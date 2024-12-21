package bundle

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/molecule"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunRest(m filesystem.Manager) *cobra.Command {
	var (
		module string
	)

	cmd := &cobra.Command{
		Use:   "rest",
		Short: "Generates a new rest molecule",
		Run: func(cmd *cobra.Command, args []string) {
			if err := parseRestInput(module); err != nil {
				panic(err)
			}

			molecule.PerformRest(m, module)
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")

	return cmd
}

func parseRestInput(module string) error {
	inputs := []input.Arg{
		{
			FieldName:  "module",
			IsRequired: true,
			Value:      module,
			Type:       input.StringType,
		},
	}

	return input.Validate(inputs)
}
