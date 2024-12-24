package unit

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/architecture/unit"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunRepo(m filesystem.Manager) *cobra.Command {
	var (
		module    string
		modelName string
	)

	cmd := &cobra.Command{
		Use:   "repo",
		Short: "Generates a new repository contract",
		Long:  "This command generates a new repository contract for the specified module and model.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := parseRepoInput(module, modelName); err != nil {
				panic(err)
			}

			repository := unit.MakeRepository(m, module, modelName)

			if err := m.GenerateFile(repository); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&modelName, "model", "n", "", "model to be managed by the repository")

	return cmd
}

func parseRepoInput(module, modelName string) error {
	inputs := []input.Arg{
		{
			FieldName:  "module",
			IsRequired: true,
			Value:      module,
			Type:       input.StringType,
		},
		{
			FieldName:  "model name",
			IsRequired: true,
			Value:      modelName,
			Type:       input.StringType,
		},
	}

	return input.Validate(inputs)
}
