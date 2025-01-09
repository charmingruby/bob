package postgres

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/shared/resource/database/postgres"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunRepo(m filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repo",
		Short: "Generates a new repository",
		Long:  "This command generates a new repository for PostgreSQL, allowing you to interact with the database using the specified module, model, and table.",
		Run: func(cmd *cobra.Command, args []string) {
			questions := []*survey.Question{
				{
					Name:     "Module",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("Module")},
					Validate: survey.Required,
				},
				{
					Name:     "ModelName",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("Model")},
					Validate: survey.Required,
				},
				{
					Name:   "TableName",
					Prompt: &survey.Input{Message: input.EnterValueMessage("Table name")},
				},
				{
					Name:     "NeedDependencies",
					Prompt:   &survey.Confirm{Message: input.EnterValueMessage("Install dependencies"), Default: true},
					Validate: survey.Required,
				},
			}

			answers := struct {
				Module           string
				ModelName        string
				TableName        string
				NeedDependencies bool
			}{}

			survey.Ask(questions, &answers)

			components, err := postgres.PerformRepository(m, answers.Module, answers.ModelName, answers.TableName, answers.NeedDependencies)
			if err != nil {
				output.ShutdownWithError(err.Error())
			}

			for _, c := range components {
				output.ComponentCreated(c.Identifier)
			}

			output.CommandSuccess("postgres repository")
		},
	}

	return cmd
}
