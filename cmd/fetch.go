package cmd

import (
	"github.com/de1ux/gitstuff/audit"
	"github.com/de1ux/gitstuff/git"
	"github.com/de1ux/gitstuff/shell"
	"github.com/spf13/cobra"
)

var FetchCmd = &cobra.Command{
	Use:  "fetch",
	Args: cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		if len(args) > 0 {
			branch := args[0]
			err = audit.Write(branch, "fetching branch "+branch+" from origin")
			if err != nil {
				return err
			}

			return shell.Spinner("> git fetch origin "+branch+":"+branch, func() error {
				return git.FetchBranch(branch)
			})
		} else {
			err = audit.Write("", "fetching all branches from all remotes")
			if err != nil {
				return err
			}

			return shell.Spinner("> git fetch --all --prune", func() error {
				return git.FetchAll()
			})
		}
	},
}