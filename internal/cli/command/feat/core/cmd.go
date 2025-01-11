package core

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/command/feat/core/bundle"
	"github.com/charmingruby/bob/internal/cli/command/feat/core/unit"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/shared/definition/component"
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

			var templateChoice string
			prompt := &survey.Select{
				Message: input.ChooseSectionMessage(section),
				Options: []string{component.UNIT_SIZE, component.BUNDLE_SIZE},
			}
			survey.AskOne(prompt, &templateChoice)

			switch templateChoice {
			case component.UNIT_SIZE:
				unit.SetupCmd(fs).Execute()
			case component.BUNDLE_SIZE:
				bundle.SetupCmd(fs).Execute()
			}
		},
	}

	return cmd
}
