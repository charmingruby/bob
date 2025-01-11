package resource

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/command/feat/resource/postgres"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "resource",
		Aliases: []string{"rsc"},
		Short:   "Generates resources (aliases: rsc)",
		Long:    "This command provides various infrastructure resource generation capabilities.",
		Run: func(cmd *cobra.Command, args []string) {
			section := "context"

			postgresOptionName := "PostgreSQL"

			var templateChoice string
			prompt := &survey.Select{
				Message: input.ChooseSectionMessage(section),
				Options: []string{postgresOptionName},
			}
			survey.AskOne(prompt, &templateChoice)

			switch templateChoice {
			case postgresOptionName:
				postgres.SetupCmd(fs).Execute()
			}
		},
	}

	cmd.AddCommand(
		postgres.SetupCmd(fs),
	)

	return cmd
}
