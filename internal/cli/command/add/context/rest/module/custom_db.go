package module

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/context/rest/module/custom_db"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunCustomDB(m filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "custom-db",
		Aliases: []string{"c-db"},
		Short:   "Generates a module with custom database (aliases: c-db)",
		Long:    "This command generates a module with a custom database, allowing you to specify the implementation.",
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
					Name:     "DatabaseName",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("Database name for persistence")},
					Validate: survey.Required,
				},
			}

			answers := struct {
				Module       string
				ModelName    string
				DatabaseName string
			}{}

			survey.Ask(questions, &answers)

			components, err := custom_db.Perform(m, answers.Module, answers.ModelName, answers.DatabaseName)
			if err != nil {
				output.ShutdownWithError(err.Error())
			}

			for _, c := range components {
				output.ComponentCreated(c.Identifier)
			}

			output.CommandSuccess("custom-db module")
		},
	}

	return cmd
}
