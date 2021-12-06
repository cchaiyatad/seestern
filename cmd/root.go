package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "seestern",
	Short: "Test Data Generator for Document-Oriented Database",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SuggestionsMinimumDistance = 1
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
}
