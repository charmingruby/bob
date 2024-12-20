package postgres

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/resource"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/spf13/cobra"
)

func RunRepository(m filesystem.Manager) *cobra.Command {
	var (
		module           string
		modelName        string
		tableName        string
		needDependencies bool
	)

	cmd := &cobra.Command{
		Use:   "repo",
		Short: "Generates a new repository",
		Run: func(cmd *cobra.Command, args []string) {
			if err := ValidateRepositoryCommandInput(module, modelName); err != nil {
				panic(err)
			}

			resource.MakeAndRunPostgresRepository(m, module, modelName, tableName, needDependencies)
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&modelName, "model", "n", "", "model to be managed by the repository")
	cmd.Flags().StringVarP(&tableName, "table name", "t", "", "table name on migrations, by default, if it is not set, it will be not created")
	cmd.Flags().BoolVarP(&needDependencies, "dependencies", "d", false, "generate dependencies")

	return cmd
}

func ValidateRepositoryCommandInput(module, name string) error {
	args := []input.Arg{
		{
			FieldName:  "module",
			Value:      module,
			IsRequired: true,
		},
		{
			FieldName:  "name",
			Value:      name,
			IsRequired: true,
		},
	}

	if err := input.ValidateArgsList(args); err != nil {
		return err
	}

	return nil
}
