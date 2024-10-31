package cmd

import (
	"strings"

	"github.com/de1ux/gitstuff/git"
	"github.com/de1ux/gitstuff/shell"
	_ "github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var (
	undoAndForce = false
	force        = false
)

func init() {
	CommitCmd.Flags().BoolVarP(&undoAndForce, "undo-and-force", "u", false, "reset the last commit, and force push the changes")
	CommitCmd.Flags().BoolVarP(&force, "force", "f", false, "force push the changes")
}

var CommitCmd = &cobra.Command{
	Use:  "commit",
	Args: cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		branch, err := git.CurrentBranch()
		if err != nil {
			return err
		}

		f := undoAndForce || force
		msg := "> git push origin " + branch
		if f {
			msg = "> git push origin -f " + branch
		}

		if git.InMergeConflict() {
			err := git.CommitNoEdit()
			if err != nil {
				return err
			}
			return shell.Spinner(msg, func() error {
				return git.Push(branch, f)
			})
		}

		commitMsg := strings.Join(args, " ")

		if strings.HasPrefix(branch, branchPrefix) {
			ticketNumber := git.TicketNumber(branchPrefix, branch)
			// prepend the ticket number
			commitMsg = ticketNumber + " " + commitMsg
		}

		if undoAndForce {
			// TODO - add a spinner
			err = git.ResetLastCommit()
			if err != nil {
				return err
			}
		}

		err = shell.Spinner("> git commit -am '"+commitMsg+"'", func() error {
			return git.Commit(commitMsg)
		})
		if err != nil {
			return err
		}

		return shell.Spinner(msg, func() error {
			return git.Push(branch, f)
		})
	},
}
