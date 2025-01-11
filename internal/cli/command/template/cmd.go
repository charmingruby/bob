package template

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/command/template/rest"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func Setup(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "template",
		Aliases: []string{"tpl"},
		Short:   "Generates project from templates (aliases: tmpl)",
		Long:    "This command creates a new project by setting up the project structure from templates.",
		Run: func(cmd *cobra.Command, args []string) {
			section := "template options"

			restOptionName := "REST API"

			var templateChoice string
			prompt := &survey.Select{
				Message: input.ChooseSectionMessage(section),
				Options: []string{restOptionName},
			}
			survey.AskOne(prompt, &templateChoice)

			switch templateChoice {
			case restOptionName:
				rest.SetupCmd(fs).Execute()
			default:
				output.ComingSoon(templateChoice, section)
			}
		},
	}

	return cmd
}
