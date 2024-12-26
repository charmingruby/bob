package module

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/context/rest/module/custom_db"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunCustomDB(m filesystem.Manager) *cobra.Command {
	var (
		module    string
		modelName string
		database  string
	)

	cmd := &cobra.Command{
		Use:     "custom-db",
		Aliases: []string{"c-db"},
		Short:   "Generates a module with custom database (aliases: c-db)",
		Long:    "This command generates a module with a custom database, allowing you to specify the implementation.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := parseCustomDBInput(module, modelName, database); err != nil {
				output.ShutdownWithError(err.Error())
			}

			components, err := custom_db.Perform(m, module, modelName, database)
			if err != nil {
				output.ShutdownWithError(err.Error())
			}

			for _, c := range components {
				output.ComponentCreated(c.Identifier)
			}

			output.CommandSuccess("custom-db module")
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module")
	cmd.Flags().StringVarP(&modelName, "modelName", "n", "", "model name")
	cmd.Flags().StringVarP(&database, "database", "d", "", "base database name to be created as example")

	return cmd
}

func parseCustomDBInput(module, model, database string) error {
	args := []input.Arg{
		{
			FieldName:  "module",
			Value:      module,
			IsRequired: true,
		},
		{
			FieldName:  "model",
			Value:      model,
			IsRequired: true,
		},
		{
			FieldName:  "database",
			Value:      database,
			IsRequired: true,
		},
	}

	return input.Validate(args)
}
