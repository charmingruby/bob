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
		Short:   "Generates bundles (aliases: bd)",
		Long:    "This command generates various units, which are individual business rules and components.",
		Run: func(cmd *cobra.Command, args []string) {
			section := "structure"

			coreName := "Model"
			serviceName := "Service with entrypoint"

			var templateChoice string
			prompt := &survey.Select{
				Message: input.ChooseSectionMessage(section),
				Options: []string{coreName, serviceName},
			}

			survey.AskOne(prompt, &templateChoice)

			switch templateChoice {
			case coreName:
				RunCore(fs).Execute()
			case serviceName:
				RunService(fs).Execute()
			}
		},
	}

	return cmd
}
