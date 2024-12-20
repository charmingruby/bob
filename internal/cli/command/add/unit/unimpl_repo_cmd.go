package unit

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/atom"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/spf13/cobra"
)

func RunUnimplementedRepository(m filesystem.Manager) *cobra.Command {
	var (
		module   string
		name     string
		database string
	)

	cmd := &cobra.Command{
		Use:   "unimpl-repo",
		Short: "Generates a new unimplemented repository",
		Run: func(cmd *cobra.Command, args []string) {
			if err := ValidateUnimplementedRepositoryCommandInput(module, name, database); err != nil {
				panic(err)
			}

			repository := atom.MakeUnimplementedRepository(m, module, name, database)
			if err := m.GenerateFile(repository); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&name, "name", "n", "", "model to be managed by the repository")
	cmd.Flags().StringVarP(&database, "database", "d", "", "database for the repository")

	return cmd
}

func ValidateUnimplementedRepositoryCommandInput(module, name, database string) error {
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
