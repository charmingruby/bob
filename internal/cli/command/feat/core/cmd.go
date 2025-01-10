package core

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/command/feat/core/bundle"
	"github.com/charmingruby/bob/internal/cli/command/feat/core/unit"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "core",
		Aliases: []string{"cr"},
		Short:   "Generates core components (aliases: cr)",
		Long:    "This command generates various module sets where the business rules are defined.",
		Run: func(cmd *cobra.Command, args []string) {
			section := "component size"

			unitName := "Unit components"
			bundleName := "Bundle components"

			var templateChoice string
			prompt := &survey.Select{
				Message: input.ChooseSectionMessage(section),
				Options: []string{unitName, bundleName},
			}
			survey.AskOne(prompt, &templateChoice)

			switch templateChoice {
			case unitName:
				unit.SetupCmd(fs).Execute()
			case bundleName:
				bundle.SetupCmd(fs).Execute()
			}
		},
	}

	return cmd
}
