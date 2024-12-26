package module

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/context/rest/template/base"
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
		Long:  "This command generates a base module, which includes the basic structure and components needed for a new module.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := parseBaseInput(module, baseModelName); err != nil {
				output.ShutdownWithError(err.Error())
			}

			if err := base.PerfomBase(m, module, baseModelName); err != nil {
				output.ShutdownWithError(err.Error())
			}

			output.CommandSuccess("base module")
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
