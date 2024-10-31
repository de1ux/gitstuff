package main

import (
	"log"

	_ "github.com/charmbracelet/bubbletea"
	"github.com/de1ux/gitstuff/cmd"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
