package unit

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/core/unit"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunRepo(m filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repo",
		Short: "Generates a new repository contract",
		Long:  "This command generates a new repository contract for the specified module and model.",
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

			repository := unit.MakeRepository(m, answers.Module, answers.ModelName)

			if err := m.GenerateFile(repository); err != nil {
				output.ShutdownWithError(err.Error())
			}

			output.ComponentCreated(repository.Identifier)
			output.CommandSuccess("repo unit")
		},
	}

	return cmd
}
