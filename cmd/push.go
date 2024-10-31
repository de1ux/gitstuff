package cmd

import (
	"github.com/de1ux/gitstuff/git"
	"github.com/de1ux/gitstuff/shell"
	"github.com/spf13/cobra"
)

var forcePush bool

func init() {
	PushCmd.Flags().BoolVarP(&forcePush, "force", "f", false, "force push to branch")
}

var PushCmd = &cobra.Command{
	Use:  "push",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var branch string
		var err error
		err = shell.Spinner("Getting current branch", func() error {
			branch, err = git.CurrentBranch()
			return err
		})
		if err != nil {
			return err
		}
		msg := "> git push origin " + branch
		if forcePush {
			msg = "> git push -f origin " + branch
		}
		return shell.Spinner(msg, func() error {
			return git.Push(branch, forcePush)
		})
	},
}
