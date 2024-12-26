package bundle

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/context/rest/bundle/handler"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunHandler(m filesystem.Manager) *cobra.Command {
	var (
		module     string
		actionName string
	)

	cmd := &cobra.Command{
		Use:     "handler",
		Aliases: []string{"hd"},
		Short:   "Generates a new REST handler (aliases: hd)",
		Long:    "This command generates a new REST handler for setting up a route.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := parseHandlerInput(module, actionName); err != nil {
				output.ShutdownWithError(err.Error())
			}

			components, err := handler.Perform(m, module, actionName)
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
	cmd.Flags().StringVarP(&actionName, "action", "a", "", "action name")

	return cmd
}

func parseHandlerInput(module, actionName string) error {
	inputs := []input.Arg{
		{
			FieldName:  "module",
			IsRequired: true,
			Value:      module,
			Type:       input.StringType,
		},
		{
			FieldName:  "action",
			IsRequired: true,
			Value:      actionName,
			Type:       input.StringType,
		},
	}

	return input.Validate(inputs)
}
