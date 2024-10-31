package cmd

import (
	"github.com/de1ux/gitstuff/git"
	"github.com/de1ux/gitstuff/shell"
	"github.com/de1ux/gitstuff/stack"
	"github.com/spf13/cobra"
)

var BackCmd = &cobra.Command{
	Use:  "back",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := stack.Load()
		if err != nil {
			return err
		}
		defer s.Save()

		branch, err := s.Back()
		if err != nil {
			return err
		}
		return shell.Spinner("> git checkout "+branch, func() error {
			return git.Checkout(branch)
		})
	},
}
