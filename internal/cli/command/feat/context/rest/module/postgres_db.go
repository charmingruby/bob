package module

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/context/rest/module/postgres"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunPostgresDB(m filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "postgres-db",
		Aliases: []string{"pg-db"},
		Short:   "Generates a module with PostgreSQL database (aliases: pg-db)",
		Long:    "This command generates a module with a PostgreSQL database implementation.",
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
				{
					Name:     "TableName",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("Table name")},
					Validate: survey.Required,
				},
			}

			answers := struct {
				Module    string
				ModelName string
				TableName string
			}{}

			survey.Ask(questions, &answers)

			components, err := postgres.Perform(m, answers.Module, answers.ModelName, answers.TableName)
			if err != nil {
				output.ShutdownWithError(err.Error())
			}

			for _, c := range components {
				output.ComponentCreated(c.Identifier)
			}

			output.CommandSuccess("postgres-db module")
		},
	}

	return cmd
}
