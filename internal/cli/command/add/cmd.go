package add

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/cli/command/add/context"
	"github.com/charmingruby/bob/internal/cli/command/add/core"
	"github.com/charmingruby/bob/internal/cli/command/add/resource"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

type Command struct {
	cmd *cobra.Command
	fs  filesystem.Manager
}

func New(cmd *cobra.Command, config *config.Configuration) *Command {
	return &Command{
		cmd: cmd,
		fs:  filesystem.New(config),
	}
}

func (c *Command) Setup() {
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
				core.SetupCmd(c.fs).Execute()
			case contextOptionsName:
				context.SetupCmd(c.fs).Execute()
			case resourceOptionsName:
				resource.SetupCmd(c.fs).Execute()
			}
		},
	}

	c.cmd.AddCommand(cmd)
}
