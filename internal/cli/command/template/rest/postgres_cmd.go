package rest

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/context/rest/template/postgres"
	"github.com/charmingruby/bob/internal/shared/filesystem"

	"github.com/spf13/cobra"
)

func RunPostgres(m filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pg",
		Short: "Creates a new project with PostgreSQL",
		Long:  "This command creates a new project using a PostgreSQL template.",
		Run: func(cmd *cobra.Command, args []string) {
			questions := []*survey.Question{
				{
					Name:     "GoVersion",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("Golang version"), Default: "1.23.3"},
					Validate: survey.Required,
				},
			}

			answers := struct {
				GoVersion string
			}{}

			survey.Ask(questions, &answers)

			components, err := postgres.PerformWithPostgres(m, answers.GoVersion)
			if err != nil {
				output.ShutdownWithError(err.Error())
			}

			for _, c := range components {
				output.ComponentCreated(c.Identifier)
			}

			output.CommandSuccess("postgres template")
		},
	}

	return cmd
}
