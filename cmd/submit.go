package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/de1ux/gitstuff/audit"
	"github.com/de1ux/gitstuff/git"
	"github.com/de1ux/gitstuff/shell"
	"github.com/spf13/cobra"
)

func init() {
	SubmitCmd.Flags().StringVar(&repo, "repo", repo, "Github repo to submit the PR to")
	SubmitCmd.Flags().StringVar(&org, "org", org, "Github organization to submit the PR to")
}

const NewPrTemplate = `
## Motivation

- [ ] I've talked about why I made this change.

<!-- Why do I want to merge this into main? -->

## Changes

- [ ] I've summarized what the new behavior should be.

<!-- Before these changes our tool would ABC. With these changes it now does XYZ. -->

## Testing

- [ ] I've written about how I've tested this.

<!-- I checked by hand / wrote an automated test -->
<!-- An existing test exercises my changes -->
<!-- I can't test this -->
`

var SubmitCmd = &cobra.Command{
	Use:  "submit",
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// If a commit message is provided, commit and push first
		if len(args) == 1 {
			commitMsg := args[0]
			if strings.HasPrefix(branch, branchPrefix) {
				ticketNumber := git.TicketNumber(branchPrefix, branch)
				commitMsg = ticketNumber + " " + commitMsg
			}

			msg := "> git commit -am '" + commitMsg + "'"
			err := audit.Write(branch, msg)
			if err != nil {
				return err
			}

			err = shell.Spinner(msg, func() error {
				return git.Commit(commitMsg)
			})
			if err != nil {
				return err
			}

			err = shell.Spinner("> git push origin "+branch, func() error {
				return git.Push(branch, false)
			})
			if err != nil {
				return err
			}
		}

		// Check if gh CLI is installed
		if err := shell.ExecOutputVerbose("gh --version"); err != nil {
			return fmt.Errorf("gh CLI is not installed. Please install it from https://cli.github.com/")
		}

		ticket := git.TicketNumber(branchPrefix, branch)
		if strings.TrimSpace(ticket) == "" {
			ticket = "Todo"
		}

		// Create a temporary file for the PR template
		tmpFile, err := os.CreateTemp("", "pr-template-*.md")
		if err != nil {
			return fmt.Errorf("failed to create temporary file: %w", err)
		}
		defer os.Remove(tmpFile.Name())

		if _, err := tmpFile.WriteString(NewPrTemplate); err != nil {
			return fmt.Errorf("failed to write template to temporary file: %w", err)
		}
		if err := tmpFile.Close(); err != nil {
			return fmt.Errorf("failed to close temporary file: %w", err)
		}

		var prURL string
		err = shell.Spinner("Creating draft PR on Github", func() error {
			// Construct the gh pr create command with all options
			ghCmd := fmt.Sprintf(
				"gh pr create --draft --title %s --head %s --base main --body-file %s",
				ticket,
				branch,
				tmpFile.Name(),
			)

			// If org and repo are specified, add them to the command
			if org != "" && repo != "" {
				ghCmd += fmt.Sprintf(" --repo %s/%s", org, repo)
			}

			// Execute the command and capture the output (PR URL)
			output, err := shell.ExecOutput(ghCmd)
			if err != nil {
				return fmt.Errorf("failed to create PR: %w", err)
			}
			prURL = strings.TrimSpace(output)
			return nil
		})
		if err != nil {
			return err
		}

		println()
		println("=========================================")
		println(prURL)
		println("=========================================")

		shell.BashStatusCode("open " + prURL)

		return audit.Write(branch, "created draft PR "+prURL)
	},
}
