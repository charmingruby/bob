package bundle

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/context/rest/bundle/setup"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunSetup(m filesystem.Manager) *cobra.Command {
	var (
		module string
	)

	cmd := &cobra.Command{
		Use:     "setup",
		Aliases: []string{"sup"},
		Short:   "Generates a new REST bundle (aliases: sup)",
		Long:    "This command generates a new REST bundle, which includes the necessary components for a RESTful API.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := parseSetupInput(module); err != nil {
				output.ShutdownWithError(err.Error())
			}

			components, err := setup.Perform(m, module)
			if err != nil {
				output.ShutdownWithError(err.Error())
			}

			for _, c := range components {
				output.ComponentCreated(c.Identifier)
			}

			output.CommandSuccess("REST bundle")
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module")

	return cmd
}

func parseSetupInput(module string) error {
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
