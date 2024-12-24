package postgres

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/resource"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunRepo(m filesystem.Manager) *cobra.Command {
	var (
		module           string
		modelName        string
		tableName        string
		needDependencies string
	)

	cmd := &cobra.Command{
		Use:   "repo",
		Short: "Generates a new repository",
		Long:  "This command generates a new repository for PostgreSQL, allowing you to interact with the database using the specified module, model, and table.",
		Run: func(cmd *cobra.Command, args []string) {
			parsedNeedDependencies, err := parseRepoInput(module, modelName, tableName, needDependencies)
			if err != nil {
				panic(err)
			}

			resource.PerformPostgresRepository(m, module, modelName, tableName, parsedNeedDependencies)
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module")
	cmd.Flags().StringVarP(&modelName, "model", "n", "", "model to be managed by the repository")
	cmd.Flags().StringVarP(&tableName, "table name", "t", "", "table name on migrations, by default, if it is not set, it will be not created")
	cmd.Flags().StringVarP(&needDependencies, "dependencies", "d", "false", "generate dependencies")

	return cmd
}

func parseRepoInput(module, modelName, tableName, needDependencies string) (bool, error) {
	needDependenciesArg := input.Arg{
		FieldName:  "dependencies",
		Value:      needDependencies,
		IsRequired: false,
		Type:       input.BoolType,
	}

	args := []input.Arg{
		{
			FieldName:  "module",
			Value:      module,
			IsRequired: true,
		},
		{
			FieldName:  "name",
			Value:      modelName,
			IsRequired: true,
		},
		{
			FieldName:  "name",
			Value:      tableName,
			IsRequired: true,
		},
		needDependenciesArg,
	}

	if err := input.Validate(args); err != nil {
		return false, err
	}

	parsedNeedDependencies, _ := input.ParseBool(needDependenciesArg)

	return parsedNeedDependencies, nil
}
