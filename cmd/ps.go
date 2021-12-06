package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// TODO-0: first usecase
// TODO-0-1: add flag
// TODO-0-2: connect to db and get all collection

// TODO-1: db pkg
// TODO-1-1: connect to mongodb and make factory pattern for ease of add new db driver in future

var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "List collections of given string connection",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ps called")
	},
}

func init() {
	rootCmd.AddCommand(psCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// psCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// psCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
