package postgres

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/resource/database/postgres/component"
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
		Use:   "repository",
		Short: "Generates a new repository",
		Run: func(cmd *cobra.Command, args []string) {
			if err := ValidateRepositoryCommandInput(module, modelName); err != nil {
				panic(err)
			}

			repo := component.MakePostgresRepository(m, module, modelName)
			if err := m.GenerateFile(repo); err != nil {
				panic(err)
			}

			if tableName != "" {
				component.RunMigration(m, tableName)
			}

			if needDependencies {
				conn := component.MakePostgresConnection(m)
				if err := m.GenerateFile(conn); err != nil {
					panic(err)
				}
			}
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
