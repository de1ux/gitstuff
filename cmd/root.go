package cmd

import (
	"strings"
	"time"

	"github.com/de1ux/gitstuff/audit"
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

	RootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		return audit.Write(timestamp + " " + cmd.Name() + " " + strings.Join(args, " "))
	}
}

var RootCmd = cobra.Command{
	Use:   "gitstuff",
	Short: "An opinionated tool for daily git tasks",
}
