package cmd

import (
	"fmt"

	"github.com/cchaiyatad/seestern/pkg/db"
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

var verbose bool
var collections []string

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringP(connectionStringKey, "s", "", "connection string to database")
	initCmd.Flags().StringP(outputKey, "o", "", "write output to <file>")

	initCmd.Flags().StringSliceVarP(&collections, collectionKey, "c", []string{}, "specific database and collection to create (in <database>.<collection> format)")
	initCmd.Flags().BoolVarP(&verbose, verboseKey, "v", false, "verbose output")

	initCmd.MarkFlagRequired(connectionStringKey)
	initCmd.MarkFlagRequired(collectionKey)

}

func initFunc(cmd *cobra.Command, args []string) {
	connectionStr := cmd.Flag(connectionStringKey).Value.String()
	out := cmd.Flag(outputKey).Value.String()

	fmt.Printf("init with %s connection string\n", connectionStr)

	param := &db.InitParam{
		CntStr:   connectionStr,
		TargetDB: collections,
		Verbose:  verbose,
	}
	fmt.Println(out)
	fmt.Println(param)
	fmt.Println(collections)
}
