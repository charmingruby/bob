package module

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/context/rest/module/postgres"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunPostgresDB(m filesystem.Manager) *cobra.Command {
	var (
		module    string
		modelName string
		tableName string
	)

	cmd := &cobra.Command{
		Use:     "postgres-db",
		Aliases: []string{"pg-db"},
		Short:   "Generates a module with PostgreSQL database (aliases: pg-db)",
		Long:    "This command generates a module with a PostgreSQL database implementation.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := parsePostgresDBInput(module, modelName, tableName); err != nil {
				output.ShutdownWithError(err.Error())
			}

			components, err := postgres.Perform(m, module, modelName, tableName)
			if err != nil {
				output.ShutdownWithError(err.Error())
			}

			for _, c := range components {
				output.ComponentCreated(c.Identifier)
			}

			output.CommandSuccess("postgres-db module")
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module")
	cmd.Flags().StringVarP(&modelName, "modelName", "n", "", "base model name")
	cmd.Flags().StringVarP(&tableName, "tableName", "t", "", "base table name")

	return cmd
}

func parsePostgresDBInput(module, model, tableName string) error {
	args := []input.Arg{
		{
			FieldName:  "module",
			Value:      module,
			IsRequired: true,
		},
		{
			FieldName:  "model name",
			Value:      model,
			IsRequired: true,
		},
		{
			FieldName:  "table name",
			Value:      tableName,
			IsRequired: true,
		},
	}

	return input.Validate(args)
}
