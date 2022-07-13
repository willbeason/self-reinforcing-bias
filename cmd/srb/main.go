package main

import (
	"os"

	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	return &cobra.Command{}
}

func main() {
	err := Cmd().Execute()
	if err != nil {
		os.Exit(1)
	}
}
