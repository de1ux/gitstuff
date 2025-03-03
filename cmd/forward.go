package cmd

import (
	"github.com/de1ux/gitstuff/audit"
	"github.com/de1ux/gitstuff/git"
	"github.com/de1ux/gitstuff/shell"
	"github.com/de1ux/gitstuff/stack"
	"github.com/spf13/cobra"
)

var ForwardCmd = &cobra.Command{
	Use:  "forward",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := stack.Load()
		if err != nil {
			return err
		}
		defer s.Save()

		branch, err := s.Forward()
		if err != nil {
			return err
		}

		err = audit.Write(branch, "checking out branch")
		if err != nil {
			return err
		}

		return shell.Spinner("> git checkout "+branch, func() error {
			return git.Checkout(branch)
		})
	},
}
