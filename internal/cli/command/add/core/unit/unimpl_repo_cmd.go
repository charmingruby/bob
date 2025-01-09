package unit

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/core/unit"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunUnimplRepo(m filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "unimpl-repo",
		Aliases: []string{"u-repo"},
		Short:   "Generates a new unimplemented repository (aliases: u-repo)",
		Long:    "This command generates a new unimplemented repository for the specified module, model, and database.",
		Run: func(cmd *cobra.Command, args []string) {
			questions := []*survey.Question{
				{
					Name:     "Module",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("Module")},
					Validate: survey.Required,
				},
				{
					Name:     "ModelName",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("Model to be managed name")},
					Validate: survey.Required,
				},
				{
					Name:     "DatabaseName",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("Database name")},
					Validate: survey.Required,
				},
			}

			answers := struct {
				Module       string
				ModelName    string
				DatabaseName string
			}{}

			survey.Ask(questions, &answers)

			repository := unit.MakeUnimplementedRepository(m, answers.Module, answers.ModelName, answers.DatabaseName)

			if err := m.GenerateFile(repository); err != nil {
				output.ShutdownWithError(err.Error())
			}

			output.ComponentCreated(repository.Identifier)
			output.CommandSuccess("unimpl-repo unit")
		},
	}

	return cmd
}
