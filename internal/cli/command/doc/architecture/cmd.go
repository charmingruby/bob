package architecture

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/shared/definition/architecture"
	"github.com/spf13/cobra"
)

func SetupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "architecture",
		Aliases: []string{"arch"},
		Short:   "Presents the overall architecture (aliases: arch)",
		Long:    "This command presents the overall architecture of the project.",
		Run: func(cmd *cobra.Command, args []string) {
			section := "architecture"

			restArchName := "REST"

			var templateChoice string
			prompt := &survey.Select{
				Message: input.ChooseSectionMessage(section),
				Options: []string{restArchName},
			}
			survey.AskOne(prompt, &templateChoice)

			switch templateChoice {
			case restArchName:
				fmt.Println(architecture.REST_ARCHITECTURE_PROPOSAL)
			}
		},
	}

	return cmd
}
