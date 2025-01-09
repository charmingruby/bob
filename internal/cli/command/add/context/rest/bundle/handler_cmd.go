package bundle

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/context/rest/bundle/handler"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunHandler(m filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "handler",
		Aliases: []string{"hd"},
		Short:   "Generates a new REST handler (aliases: hd)",
		Long:    "This command generates a new REST handler for setting up a route.",
		Run: func(cmd *cobra.Command, args []string) {
			questions := []*survey.Question{
				{
					Name:     "Module",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("Module")},
					Validate: survey.Required,
				},
				{
					Name:     "HandlerName",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("Handler name")},
					Validate: survey.Required,
				},
			}

			answers := struct {
				Module      string
				HandlerName string
			}{}

			survey.Ask(questions, &answers)

			components, err := handler.Perform(m, answers.Module, answers.HandlerName)
			if err != nil {
				output.ShutdownWithError(err.Error())
			}

			for _, c := range components {
				output.ComponentCreated(c.Identifier)
			}

			output.CommandSuccess("REST bundle")
		},
	}

	return cmd
}
