package bundle

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "bundle",
		Aliases: []string{"bd"},
		Short:   "Generates module sets (aliases: bd)",
		Long:    "This command generates various bundles, which are simpler sets of components.",
		Run: func(cmd *cobra.Command, args []string) {
			section := "structure"

			handlerName := "New handler"
			setupName := "Setup context components"

			var templateChoice string
			prompt := &survey.Select{
				Message: input.ChooseSectionMessage(section),
				Options: []string{handlerName, setupName},
			}

			survey.AskOne(prompt, &templateChoice)

			switch templateChoice {
			case handlerName:
				RunHandler(fs).Execute()
			case setupName:
				RunSetup(fs).Execute()
			}
		},
	}

	return cmd
}
