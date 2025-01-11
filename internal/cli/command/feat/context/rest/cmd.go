package rest

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/command/feat/context/rest/bundle"
	"github.com/charmingruby/bob/internal/cli/command/feat/context/rest/module"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/shared/definition/component"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "rest",
		Aliases: []string{"rs"},
		Short:   "Generates REST components (aliases: rs)",
		Long:    "This command generates REST components, which are sets of base o REST APIs.",
		Run: func(cmd *cobra.Command, args []string) {
			section := "component size"

			var templateChoice string
			prompt := &survey.Select{
				Message: input.ChooseSectionMessage(section),
				Options: []string{component.BUNDLE_SIZE, component.MODULE_SIZE},
			}

			survey.AskOne(prompt, &templateChoice)

			switch templateChoice {
			case component.BUNDLE_SIZE:
				module.SetupCmd(fs).Execute()
			case component.MODULE_SIZE:
				bundle.SetupCmd(fs).Execute()
			}
		}}

	return cmd
}
