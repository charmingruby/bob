package unit

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/core/unit"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunModel(m filesystem.Manager) *cobra.Command {
	var (
		module    string
		modelName string
	)

	cmd := &cobra.Command{
		Use:     "model",
		Aliases: []string{"mdl"},
		Short:   "Generates a new model (aliases: mdl)",
		Long:    "This command generates a new model to me managed and persisted.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := parseModelInput(module, modelName); err != nil {
				output.ShutdownWithError(err.Error())
			}

			model := unit.MakeModel(m, module, modelName)

			if err := m.GenerateFile(model); err != nil {
				output.ShutdownWithError(err.Error())
			}

			output.ComponentCreated(model.Identifier)
			output.CommandSuccess("model unit")
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module")
	cmd.Flags().StringVarP(&modelName, "name", "n", "", "model name")

	return cmd
}

func parseModelInput(module, modelName string) error {
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
