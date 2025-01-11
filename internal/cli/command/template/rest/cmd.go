package rest

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "rest",
		Aliases: []string{"rs"},
		Short:   "Generates REST API project from templates (aliases: rs)",
		Long:    "This command provides various template resources to generate a REST API project.",
		Run: func(cmd *cobra.Command, args []string) {
			section := "template options"

			baseTemplateName := "Base - Without resources implementations"
			postgresTemplateName := "PostgreSQL - Contains PostgreSQL implementations"

			var templateChoice string
			prompt := &survey.Select{
				Message: input.ChooseSectionMessage(section),
				Options: []string{baseTemplateName, postgresTemplateName},
			}
			survey.AskOne(prompt, &templateChoice)

			switch templateChoice {
			case baseTemplateName:
				RunBase(fs).Execute()
			case postgresTemplateName:
				RunPostgres(fs).Execute()
			}
		},
	}

	return cmd
}
