package cmd

import (
	"fmt"

	"github.com/cchaiyatad/seestern/pkg/db"
	"github.com/spf13/cobra"
)

const (
	connectionStringKey = "connectionString"
	databaseKey         = "database"
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
	psCmd.MarkFlagRequired(connectionStringKey)
}

func ps(cmd *cobra.Command, args []string) {
	connectionStr := cmd.Flag(connectionStringKey).Value.String()
	database := cmd.Flag(databaseKey).Value.String()

	fmt.Printf("list collections form %s connection string\n", connectionStr)

	param := &db.PSParam{
		CntStr: connectionStr,
		Vendor: "mongo",
		DBName: database,
	}

	info, err := db.PS(param)
	cobra.CheckErr(err)
	fmt.Print(info)
}
