package cmd

import (
	"context"
	"os"
	"strings"

	"github.com/de1ux/gitstuff/git"
	"github.com/de1ux/gitstuff/shell"
	"github.com/google/go-github/v50/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

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
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := github.NewClient(
			oauth2.NewClient(
				context.Background(),
				oauth2.StaticTokenSource(
					&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
				),
			),
		)

		current, err := git.CurrentBranch()
		if err != nil {
			return err
		}
		ticket := git.TicketNumber(branchPrefix, current)
		if strings.TrimSpace(ticket) == "" {
			ticket = "Todo"
		}

		var pr *github.PullRequest
		err = shell.Spinner("Creating draft PR on Github", func() error {
			pr, _, err = c.PullRequests.Create(cmd.Context(), "DataDog", "dd-source", &github.NewPullRequest{
				Title: github.String(ticket),
				Draft: github.Bool(true),
				Head:  github.String(current),
				Base:  github.String("main"),
				Body:  github.String(NewPrTemplate),
			})
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
		println()
		println("=========================================")
		println(pr.GetHTMLURL())
		println("=========================================")

		shell.BashStatusCode("open " + pr.GetHTMLURL())

		return nil
	},
}
