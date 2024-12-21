package bundle

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/molecule"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunCore(m filesystem.Manager) *cobra.Command {
	var (
		module    string
		modelName string
	)

	cmd := &cobra.Command{
		Use:   "core",
		Short: "Generates a new core molecule",
		Run: func(cmd *cobra.Command, args []string) {
			if err := parseCoreInput(module, modelName); err != nil {
				panic(err)
			}

			molecule.PerformCore(m, module, modelName)
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module")
	cmd.Flags().StringVarP(&modelName, "model", "n", "", "base model name")

	return cmd
}

func parseCoreInput(module, modelName string) error {
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
