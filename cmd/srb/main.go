package main

import (
	"github.com/spf13/cobra"
	"os"
)

var cmd = &cobra.Command{}

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
