package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	collectionKey = "collection"
	outputKey     = "output"
	verboseKey    = "verbose"
)

// TODO-2: second usecase

// TODO-3: create config file pkg

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a configuration file (.ss.toml)",
	Run:   initFunc,
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringP(connectionStringKey, "s", "", "connection string to database")
	initCmd.Flags().StringSliceP(collectionKey, "c", []string{}, "specific database and collection to create (in <database>.<collection> format)")
	initCmd.Flags().StringP(outputKey, "o", "", "write output to <file>")
	initCmd.Flags().BoolP(verboseKey, "v", false, "verbose output")

	initCmd.MarkFlagRequired(connectionStringKey)
	initCmd.MarkFlagRequired(collectionKey)

}

func initFunc(cmd *cobra.Command, args []string) {
	fmt.Println("init called")
	connectionStr := cmd.Flag(connectionStringKey).Value.String()
	out := cmd.Flag(outputKey).Value.String()
	verbose := cmd.Flag(verboseKey).Value

	collections := cmd.Flag(collectionKey).Value.String()

	fmt.Println(connectionStr)
	fmt.Println(out)
	fmt.Println(verbose)
	fmt.Println(collections)
}
