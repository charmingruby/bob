package bundle

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/context/rest/bundle/setup"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunSetup(m filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "setup",
		Aliases: []string{"sup"},
		Short:   "Generates a new REST bundle (aliases: sup)",
		Long:    "This command generates a new REST bundle, which includes the necessary components for a RESTful API.",
		Run: func(cmd *cobra.Command, args []string) {
			questions := []*survey.Question{
				{
					Name:     "Module",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("Module")},
					Validate: survey.Required,
				},
			}

			answers := struct {
				Module string
			}{}

			survey.Ask(questions, &answers)

			components, err := setup.Perform(m, answers.Module)
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
