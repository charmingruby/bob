package module

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/organism"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/spf13/cobra"
)

func RunModuleWithPostgresDatabase(m filesystem.Manager) *cobra.Command {
	var (
		module    string
		modelName string
		tableName string
	)

	cmd := &cobra.Command{
		Use:   "postgres-db",
		Short: "Generates a module with postgres database",
		Run: func(cmd *cobra.Command, args []string) {
			if err := input.ValidateOnlyModuleCommandInput(module); err != nil {
				panic(err)
			}

			organism.PerformModuleWithPostgresDatabase(m, module, modelName, tableName)
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&modelName, "modelName", "n", "", "model name")
	cmd.Flags().StringVarP(&tableName, "tableName", "t", "", "tableName name")

	return cmd
}

func ValidateModuleWithPostgresDatabaseCommandInput(module, model, tableName string) error {
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
			FieldName:  "table name",
			Value:      tableName,
			IsRequired: true,
		},
	}

	if err := input.ValidateArgsList(args); err != nil {
		return err
	}

	return nil
}
