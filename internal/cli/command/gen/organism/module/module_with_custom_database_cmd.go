package module

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/organism"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/spf13/cobra"
)

func RunModuleWithCustomDatabase(m filesystem.Manager) *cobra.Command {
	var (
		module    string
		modelName string
		database  string
	)

	cmd := &cobra.Command{
		Use:   "w-db",
		Short: "Generates a module with custom database",
		Run: func(cmd *cobra.Command, args []string) {
			if err := ValidateModuleWithCustomDatabaseCommandInput(module, modelName, database); err != nil {
				panic(err)
			}

			organism.MakeAndRunModuleWithCustomDatabase(m, module, modelName, database)
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&modelName, "modelName", "n", "", "model name")
	cmd.Flags().StringVarP(&database, "database", "d", "", "database name")

	return cmd
}

func ValidateModuleWithCustomDatabaseCommandInput(module, model, database string) error {
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

	if err := input.ValidateArgsList(args); err != nil {
		return err
	}

	return nil
}
