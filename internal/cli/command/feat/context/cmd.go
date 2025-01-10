package context

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/command/feat/context/rest"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "context",
		Aliases: []string{"ctx"},
		Short:   "Generates components from determinated contexts (aliases: ctx)",
		Long:    "This command creates commands from determinated contexts.",
		Run: func(cmd *cobra.Command, args []string) {
			section := "context"

			restOptionName := "REST API"
			gRPCOptionName := "gRPC API"

			var templateChoice string
			prompt := &survey.Select{
				Message: input.ChooseSectionMessage(section),
				Options: []string{restOptionName, gRPCOptionName},
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
