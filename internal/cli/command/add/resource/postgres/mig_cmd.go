package postgres

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/resource"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunMig(m filesystem.Manager) *cobra.Command {
	var (
		tableName string
	)

	cmd := &cobra.Command{
		Use:   "mig",
		Short: "Generates a new PostgreSQL migration",
		Long:  "This command generates a new migration file for PostgreSQL, allowing you to define changes to your database schema.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := parseMigInput(tableName); err != nil {
				output.ShutdownWithError(err.Error())
			}

			if err := resource.PerformPostgresMigration(m, tableName); err != nil {
				output.ShutdownWithError(err.Error())
			}

			output.CommandSuccess("postgres migration")
		},
	}

	cmd.Flags().StringVarP(&tableName, "table name", "t", "examples", "table name to be created")

	return cmd
}

func parseMigInput(tableName string) error {
	args := []input.Arg{
		{
			FieldName:  "table name",
			Value:      tableName,
			IsRequired: true,
		},
	}

	return input.Validate(args)
}
