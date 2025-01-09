package module

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/context/rest/module/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunBase(m filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "base",
		Short: "Generates a base module",
		Long:  "This command generates a base module, which includes the basic structure and components needed for a new module.",
		Run: func(cmd *cobra.Command, args []string) {
			questions := []*survey.Question{
				{
					Name:     "Module",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("Module")},
					Validate: survey.Required,
				},
				{
					Name:     "ModelName",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("Base model name")},
					Validate: survey.Required,
				},
			}

			answers := struct {
				Module    string
				ModelName string
			}{}

			survey.Ask(questions, &answers)

			components, err := base.Perform(m, answers.Module, answers.ModelName)
			if err != nil {
				output.ShutdownWithError(err.Error())
			}

			for _, c := range components {
				output.ComponentCreated(c.Identifier)
			}

			output.CommandSuccess("base module")
		},
	}

	return cmd
}
