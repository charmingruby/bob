package organism

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/organism"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/spf13/cobra"
)

func RunModule(m filesystem.Manager) *cobra.Command {
	var (
		module   string
		database string
	)

	cmd := &cobra.Command{
		Use:   "mod",
		Short: "Generates a module",
		Run: func(cmd *cobra.Command, args []string) {
			if err := input.ValidateOnlyModuleCommandInput(module); err != nil {
				panic(err)
			}

			organism.MakeAndRunBaseModule(m, module, database)
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&database, "database", "d", "", "database to implement repository, if it is not set, it will be not created")

	return cmd
}

func ValidateRepositoryCommandInput(module, name, database string) error {
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
		{
			FieldName: "database",
			Value:     database,
		},
	}

	if err := input.ValidateArgsList(args); err != nil {
		return err
	}

	return nil
}
