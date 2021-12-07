package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// TODO-4: fourth usecase

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a test data from configuration file (.ss.toml)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}