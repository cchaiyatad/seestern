package cmd

import (
	"github.com/cchaiyatad/seestern/internal/log"
	"github.com/cchaiyatad/seestern/pkg/db"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a test data from configuration file",
	Run:   genFunc,
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP(connectionStringKey, "s", "", "connection string to database")
	generateCmd.Flags().StringP(fileKey, "f", "", "path to configuration file")
	generateCmd.Flags().StringP(outputKey, "o", "", "path to save generate data")

	generateCmd.Flags().BoolVarP(&verbose, verboseKey, "v", false, "verbose output")
	generateCmd.Flags().BoolVarP(&isDrop, dropKey, "d", false, "drop all document in collection in configuration file")
	generateCmd.Flags().BoolVarP(&isInsert, insertKey, "i", false, "insert document in collection in configuration file")

	generateCmd.MarkFlagRequired(fileKey)
}

func genFunc(cmd *cobra.Command, _ []string) {
	connectionStr := cmd.Flag(connectionStringKey).Value.String()
	file := cmd.Flag(fileKey).Value.String()
	out := cmd.Flag(outputKey).Value.String()

	log.Logf(log.Info, "generate with configuration file %s", file)
	param := &db.GenParam{
		CntStr:   connectionStr,
		Vendor:   "mongo",
		File:     file,
		Outpath:  out,
		Verbose:  verbose,
		IsDrop:   isDrop,
		IsInsert: isInsert,
	}

	if err := isEitherVerboseOrOutSet(out, verbose); err != nil {
		log.Log(log.Error, err)
		cobra.CheckErr(err)
	}

	if err := isCntStrSetWhenEitherDropOrInsertSet(connectionStr, isDrop, isInsert); err != nil {
		log.Log(log.Error, err)
		cobra.CheckErr(err)
	}

	err := db.Gen(param)
	if err != nil {
		log.Log(log.Error, err)
		cobra.CheckErr(err)
	}

	if out != "" {
		log.Logf(log.Info, "generated data is saved to %s\n", out)
	}
}
