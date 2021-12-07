package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const connectionStringKey = "connectionString"

// TODO-0: first usecase
// TODO-0-2: connect to db and get all collection

// TODO-1: db pkg
// TODO-1-1: connect to mongodb and make factory pattern for ease of add new db driver in future

var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "List collections of given connection string",
	Run:   ps,
}

func init() {
	rootCmd.AddCommand(psCmd)

	psCmd.Flags().StringP(connectionStringKey, "s", "", "connection string to database")
	psCmd.MarkFlagRequired(connectionStringKey)
}

func ps(cmd *cobra.Command, args []string) {
	connectionStr := cmd.Flag(connectionStringKey).Value
	fmt.Printf("list all collection form %s database\n", connectionStr)
}
