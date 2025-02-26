package cmd

import (
	"github.com/spf13/cobra"
)

var branchPrefix = ""

func init() {
	RootCmd.Flags().StringVar(&branchPrefix, "branch-prefix", "ne", "the prefix for new branches")

	RootCmd.AddCommand(InitCmd)
	RootCmd.AddCommand(CommitCmd)
	RootCmd.AddCommand(SubmitCmd)
	RootCmd.AddCommand(PullCmd)
	RootCmd.AddCommand(PushCmd)
	RootCmd.AddCommand(CheckoutCmd)
	RootCmd.AddCommand(BackCmd)
	RootCmd.AddCommand(ForwardCmd)
	RootCmd.AddCommand(OpenCmd)
}

var RootCmd = cobra.Command{
	Use:   "gitstuff",
	Short: "An opinionated tool for daily git tasks",
}
