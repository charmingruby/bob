package module

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/architecture/structure"
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
		Use:   "custom-db",
		Short: "Generates a module with custom database",
		Run: func(cmd *cobra.Command, args []string) {
			if err := parseCustomDBInput(module, modelName, database); err != nil {
				panic(err)
			}

			structure.PerformModuleWithCustomDatabase(m, module, modelName, database)
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
