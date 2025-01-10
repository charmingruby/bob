package add

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/command/add/context"
	"github.com/charmingruby/bob/internal/cli/command/add/core"
	"github.com/charmingruby/bob/internal/cli/command/add/resource"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func Setup(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Generates new components",
		Run: func(cmd *cobra.Command, args []string) {
			section := "component type"

			coreOptionName := "Core"
			contextOptionsName := "Context"
			resourceOptionsName := "Resource"

			var templateChoice string
			prompt := &survey.Select{
				Message: input.ChooseSectionMessage(section),
				Options: []string{coreOptionName, contextOptionsName, resourceOptionsName},
			}
			survey.AskOne(prompt, &templateChoice)

			switch templateChoice {
			case coreOptionName:
				core.SetupCmd(fs).Execute()
			case contextOptionsName:
				context.SetupCmd(fs).Execute()
			case resourceOptionsName:
				resource.SetupCmd(fs).Execute()
			}
		},
	}

	return cmd
}
