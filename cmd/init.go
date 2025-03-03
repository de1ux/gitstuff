package cmd

import (
	"fmt"
	"log"

	"github.com/de1ux/gitstuff/audit"
	"github.com/de1ux/gitstuff/git"
	"github.com/de1ux/gitstuff/shell"
	"github.com/spf13/cobra"
)

var (
	branchDefault = ""
	skipPull      = false
)

func init() {
	InitCmd.Flags().StringVar(&branchDefault, "default-branch", "main", "the base for new branches")
	InitCmd.Flags().BoolVar(&skipPull, "skip-pull", false, "skip pulling the default-branch on new branches")
}

var InitCmd = &cobra.Command{
	Use:  "init",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO - prompt for ticket if not supplied
		ticket := args[0]
		// TODO - what if branch prefix is "" ?
		branch := branchPrefix + "/" + ticket

		untracked, err := git.UntrackedFiles()
		if err != nil {
			return err
		}
		if untracked {
			shell.PromptExit("There are untracked files in this repo. Continue?")
		}

		isDefault, err := git.IsBranch(branchDefault)
		if err != nil {
			return err
		}
		if !isDefault {
			yes := shell.PromptYes(fmt.Sprintf("You are not on %s branch. Check it out?", branchDefault))
			if yes {
				// switch to default branch
				err = git.Checkout(branchDefault)
				if err != nil {
					return err
				}
				// pull default branch
				if !skipPull {
					err = git.Pull("")
					if err != nil {
						return err
					}
				}
			} else {
				log.Println("Branching off of " + branch + " instead of " + branchDefault)
			}
		}

		branchExists, err := git.BranchExists(branch)
		if branchExists {
			shell.PromptExit("Branch " + branch + " exists, erase and reinit?")
			err = git.Checkout(branchDefault)
			if err != nil {
				return err
			}
			err = git.DeleteBranch(branch)
			if err != nil {
				return err
			}
		}

		err = git.CheckoutNew(branch)
		if err != nil {
			return err
		}

		log.Println("All set, gl hf")

		return audit.Write(branch, "created new branch")

	},
}
