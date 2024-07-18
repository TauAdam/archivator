package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Short: "Plain archiver",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		handleError(err)
	}
}

func handleError(err error) {
	if err != nil {
		_ = fmt.Errorf("Error: %v", err)
		os.Exit(1)
	}
}
