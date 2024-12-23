package module

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/architecture/structure"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunBase(m filesystem.Manager) *cobra.Command {
	var (
		module        string
		baseModelName string
	)

	cmd := &cobra.Command{
		Use:   "base",
		Short: "Generates a base module",
		Run: func(cmd *cobra.Command, args []string) {
			if err := parseBaseInput(module, baseModelName); err != nil {
				panic(err)
			}

			structure.PerformBaseModule(m, module, baseModelName)
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module")
	cmd.Flags().StringVarP(&baseModelName, "model", "n", "", "base model name to be created as example")

	return cmd
}

func parseBaseInput(module, baseModelName string) error {
	inputs := []input.Arg{
		{
			FieldName:  "module",
			IsRequired: true,
			Value:      module,
			Type:       input.StringType,
		},
		{
			FieldName:  "base model name",
			IsRequired: true,
			Value:      baseModelName,
			Type:       input.StringType,
		},
	}

	return input.Validate(inputs)
}
