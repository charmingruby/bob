package rest

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/context/rest/template/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"

	"github.com/spf13/cobra"
)

func RunBase(m filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "base",
		Short: "Creates a new project from a base template",
		Long:  "This command creates a new project using a customizable base persistence layer.",
		Run: func(cmd *cobra.Command, args []string) {
			questions := []*survey.Question{
				{
					Name:     "GoVersion",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("Golang version"), Default: "1.23.3"},
					Validate: survey.Required,
				},
				{
					Name:     "Database",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("database name"), Help: "e.g. postgres, mysql, sqlite"},
					Validate: survey.Required,
				},
			}

			answers := struct {
				GoVersion string
				Database  string
			}{}

			survey.Ask(questions, &answers)

			components, err := base.Perfom(m, answers.GoVersion, answers.Database)
			if err != nil {
				output.ShutdownWithError(err.Error())
			}

			for _, c := range components {
				output.ComponentCreated(c.Identifier)
			}

			output.CommandSuccess("base template")
		},
	}

	return cmd
}
