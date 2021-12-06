package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// TODO-2: second usecase
// TODO-2-1: add flag

// TODO-3: create config file pkg

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a configuration file (.ss.toml)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
