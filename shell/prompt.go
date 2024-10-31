package shell

import (
	"fmt"
	"os"

	"github.com/erikgeiser/promptkit/confirmation"
	"github.com/pterm/pterm"
)

func PromptExit(msg string) {
	input := confirmation.New(msg, confirmation.Undecided)
	_, err := input.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func PromptYes(msg string) bool {
	input := confirmation.New(msg, confirmation.Undecided)
	b, err := input.RunPrompt()
	if err != nil {
		return b
	}
	return b
}

func Spinner(msg string, f func() error) error {
	spinner, err := pterm.DefaultSpinner.WithShowTimer(false).Start(msg)
	if err != nil {
		return err
	}
	err = f()
	if err != nil {
		return err
	}
	return spinner.Stop()
}
