package cmd

import (
	"fmt"

	"github.com/cchaiyatad/seestern/internal/log"
	"github.com/cchaiyatad/seestern/pkg/db"
	"github.com/spf13/cobra"
)

var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "List collections of given connection string",
	Run:   ps,
}

func init() {
	rootCmd.AddCommand(psCmd)

	psCmd.Flags().StringP(connectionStringKey, "s", "", "connection string to database")
	psCmd.Flags().StringP(databaseKey, "d", "", "specific database to list collection")

	_ = psCmd.MarkFlagRequired(connectionStringKey)
}

func ps(cmd *cobra.Command, _ []string) {
	connectionStr := cmd.Flag(connectionStringKey).Value.String()
	database := cmd.Flag(databaseKey).Value.String()

	log.Logf(log.Info, "list collections form %s connection string\n", connectionStr)

	param := &db.PSParam{
		CntStr: connectionStr,
		Vendor: "mongo",
		DBName: database,
	}

	info, err := db.PS(param)

	if err != nil {
		log.Log(log.Error, err)
		cobra.CheckErr(err)
	}

	fmt.Print(info)
}
