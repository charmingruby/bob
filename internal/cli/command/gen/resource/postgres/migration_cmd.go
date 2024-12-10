package postgres

import (
	"github.com/charmingruby/bob/internal/component/resource/database/postgres/component"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/spf13/cobra"
)

func RunMigration(m filesystem.Manager) *cobra.Command {
	var (
		tableName string
	)

	cmd := &cobra.Command{
		Use:   "migration",
		Short: "Generates a new migration",
		Run: func(cmd *cobra.Command, args []string) {
			component.RunMigration(m, tableName)
		},
	}

	cmd.Flags().StringVarP(&tableName, "table name", "t", "", "table name on migrations, by default, if it is not set, it will be not created")

	return cmd
}
