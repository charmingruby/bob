package setup

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/spf13/cobra"
)

func Setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "setup",
		Short: "Setup project configurations and updates",
		Long:  "This command allows you to setup project configurations.",
		Run: func(cmd *cobra.Command, args []string) {
			section := "command"

			initName := "Create project configuration file"
			updateName := "Update Bob version"

			var templateChoice string
			prompt := &survey.Select{
				Message: input.ChooseSectionMessage(section),
				Options: []string{initName, updateName},
			}

			survey.AskOne(prompt, &templateChoice)

			switch templateChoice {
			case initName:
				RunInit().Execute()
			case updateName:
				RunUpdate().Execute()
			}
		},
	}

	return cmd
}
