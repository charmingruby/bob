package doc

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/command/doc/architecture"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/spf13/cobra"
)

func Setup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "doc",
		Short: "Presents documentations",
		Long:  "This command presents the various documentations used for the project decisions.",
		Run: func(cmd *cobra.Command, args []string) {
			section := "document type"

			architecturalDocName := "Architecture"

			var templateChoice string
			prompt := &survey.Select{
				Message: input.ChooseSectionMessage(section),
				Options: []string{architecturalDocName},
			}

			survey.AskOne(prompt, &templateChoice)

			switch templateChoice {
			case architecturalDocName:
				architecture.SetupCmd().Execute()
			}
		},
	}

	return cmd
}
