package cmd

import (
	"github.com/de1ux/gitstuff/git"
	"github.com/spf13/cobra"
)

var branchPrefix = ""
var branch = ""
var repo = ""
var org = ""
var force = false

func init() {
	var err error
	branch, err = git.CurrentBranch()
	if err != nil {
		println("Failed to get current branch, are you in a git repo?")
		panic(err)
	}
	org, repo, err = git.CurrentOrgAndRepo()
	if err != nil {
		println("Failed to get current repo, are you in a git repo?")
		panic(err)
	}

	RootCmd.Flags().StringVar(&branchPrefix, "branch-prefix", "ne", "the prefix for new branches")

	RootCmd.AddCommand(InitCmd)
	RootCmd.AddCommand(CommitCmd)
	RootCmd.AddCommand(SubmitCmd)
	RootCmd.AddCommand(PullCmd)
	RootCmd.AddCommand(PushCmd)
	RootCmd.AddCommand(FetchCmd)
	RootCmd.AddCommand(CheckoutCmd)
	RootCmd.AddCommand(BackCmd)
	RootCmd.AddCommand(ForwardCmd)
	RootCmd.AddCommand(OpenCmd)
}

var RootCmd = cobra.Command{
	Use:   "gitstuff",
	Short: "An opinionated tool for daily git tasks",
}
