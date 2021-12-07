package cmd

import (
	"fmt"

	"github.com/cchaiyatad/seestern/pkg/db"
	"github.com/spf13/cobra"
)

const connectionStringKey = "connectionString"
const database = "database"

var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "List collections of given connection string",
	Run:   ps,
}

func init() {
	rootCmd.AddCommand(psCmd)

	psCmd.Flags().StringP(connectionStringKey, "s", "", "connection string to database")
	psCmd.Flags().StringP(database, "d", "", "specific database to list collection")
	psCmd.MarkFlagRequired(connectionStringKey)
}

func ps(cmd *cobra.Command, args []string) {
	connectionStr := cmd.Flag(connectionStringKey).Value.String()
	database := cmd.Flag(database).Value.String()

	fmt.Printf("list collections form %s connection string\n", connectionStr)

	info, err := db.PS(connectionStr, "mongo", database)
	cobra.CheckErr(err)
	fmt.Print(info)
}
