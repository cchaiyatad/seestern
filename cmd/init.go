package cmd

import (
	"github.com/cchaiyatad/seestern/internal/log"
	"github.com/cchaiyatad/seestern/pkg/db"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a configuration file (.ss.toml)",
	Run:   initFunc,
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringP(connectionStringKey, "s", "", "connection string to database")
	initCmd.Flags().StringP(outputKey, "o", "", "path to create output file")
	initCmd.Flags().StringP(fileTypeKey, "t", "yaml", "file type of output file (json, yaml, toml)")

	initCmd.Flags().StringSliceVarP(&collections, collectionKey, "c", []string{}, "specific database and collection to create (in <database>.<collection> format)")
	initCmd.Flags().BoolVarP(&verbose, verboseKey, "v", false, "verbose output")

	_ = initCmd.MarkFlagRequired(connectionStringKey)
	_ = initCmd.MarkFlagRequired(collectionKey)
}

func initFunc(cmd *cobra.Command, _ []string) {
	connectionStr := cmd.Flag(connectionStringKey).Value.String()
	out := cmd.Flag(outputKey).Value.String()
	fileType := cmd.Flag(fileTypeKey).Value.String()

	log.Logf(log.Info, "init with %s connection string\n", connectionStr)
	if err := isEitherVerboseOrOutSet(out, verbose); err != nil {
		log.Log(log.Error, err)
		cobra.CheckErr(err)
	}

	param := &db.InitParam{
		CntStr:      connectionStr,
		Vendor:      "mongo",
		TargetColls: collections,
		Outpath:     out,
		Verbose:     verbose,
		FileType:    fileType,
	}

	path, err := db.Init(param)
	if err != nil {
		log.Log(log.Error, err)
		cobra.CheckErr(err)
	}

	if out != "" {
		log.Logf(log.Info, "config file is saved to %s\n", path)
	}
}
