package cmd

import (
	"strings"

	"github.com/de1ux/gitstuff/audit"
	"github.com/de1ux/gitstuff/git"
	"github.com/de1ux/gitstuff/shell"
	"github.com/spf13/cobra"
)

var newBranch bool
var fetchBranch bool

func init() {
	CheckoutCmd.Flags().BoolVarP(&newBranch, "new", "b", false, "make a new branch and checkout to it")
	CheckoutCmd.Flags().BoolVarP(&fetchBranch, "fetch", "u", false, "fetch the branch before checking it out")
}

var CheckoutCmd = &cobra.Command{
	Use:  "checkout",
	Args: cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		branch := args[0]

		if newBranch {
			err := audit.Write(branch, "creating new branch and checking out")
			if err != nil {
				return err
			}

			return shell.Spinner("> git checkout -b "+branch, func() error {
				return git.CheckoutNew(branch)
			})
		}

		if fetchBranch {
			err := audit.Write(branch, "fetching and checking out branch")
			if err != nil {
				return err
			}

			return shell.Spinner("> git fetch origin "+branch+":"+branch+" && git checkout "+branch, func() error {
				err := git.FetchBranch(branch)
				if err != nil {
					return err
				}
				return git.Checkout(branch)
			})
		}

		err := audit.Write(branch, "checking out branch")
		if err != nil {
			return err
		}

		if len(args) == 1 {
			return shell.Spinner("> git checkout "+branch, func() error {
				return git.Checkout(branch)
			})
		}

		return shell.ExecOutputVerbose("git checkout " + strings.Join(args, " "))
	},
}
