package postgres

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/shared/resource/database/postgres"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunMig(m filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mig",
		Short: "Generates a new PostgreSQL migration",
		Long:  "This command generates a new migration file for PostgreSQL, allowing you to define changes to your database schema.",
		Run: func(cmd *cobra.Command, args []string) {
			questions := []*survey.Question{
				{
					Name:     "TableName",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("Table name")},
					Validate: survey.Required,
				},
			}

			answers := struct {
				TableName string
			}{}

			survey.Ask(questions, &answers)

			components, err := postgres.PerformMigration(m, answers.TableName)
			if err != nil {
				output.ShutdownWithError(err.Error())
			}

			for _, c := range components {
				output.ComponentCreated(c.Identifier)
			}

			output.CommandSuccess("postgres migration")
		},
	}

	return cmd
}
