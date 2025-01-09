package unit

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/core/unit"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunModel(m filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "model",
		Aliases: []string{"mdl"},
		Short:   "Generates a new model (aliases: mdl)",
		Long:    "This command generates a new model to me managed and persisted.",
		Run: func(cmd *cobra.Command, args []string) {
			questions := []*survey.Question{
				{
					Name:     "Module",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("Module")},
					Validate: survey.Required,
				},
				{
					Name:     "ModelName",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("Model name")},
					Validate: survey.Required,
				},
			}

			answers := struct {
				Module    string
				ModelName string
			}{}

			survey.Ask(questions, &answers)

			model := unit.MakeModel(m, answers.Module, answers.ModelName)

			if err := m.GenerateFile(model); err != nil {
				output.ShutdownWithError(err.Error())
			}

			output.ComponentCreated(model.Identifier)
			output.CommandSuccess("model unit")
		},
	}

	return cmd
}
