package cmd

import (
	"github.com/de1ux/gitstuff/git"
	"github.com/de1ux/gitstuff/shell"
	"github.com/spf13/cobra"
)

var PullCmd = &cobra.Command{
	Use:  "pull",
	Args: cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var branch string
		if len(args) > 0 {
			branch = args[1]
		} else {
			err = shell.Spinner("Getting current branch", func() error {
				branch, err = git.CurrentBranch()
				return err
			})
			if err != nil {
				return err
			}
		}
		return shell.Spinner("> git pull origin "+branch, func() error {
			return git.Pull(branch)
		})
	},
}
