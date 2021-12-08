package cmd

import (
	"errors"
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
	initCmd.Flags().StringP(outputKey, "o", "", "path to create output file")

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
		CntStr:      connectionStr,
		Vendor:      "mongo",
		TargetColls: collections,
		Output:      out,
		Verbose:     verbose,
	}

	cobra.CheckErr(isFlagValid(out, verbose))

	path, err := db.Init(param)
	cobra.CheckErr(err)
	if out != "" {
		fmt.Println(path)
	}
}

func isFlagValid(out string, verbose bool) error {
	if out == "" && !verbose {
		return errors.New("if verbose is not set, output has to be set")
	}
	return nil
}
