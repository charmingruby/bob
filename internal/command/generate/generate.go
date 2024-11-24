package generate

import (
	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/command/generate/brick"
	"github.com/charmingruby/bob/internal/command/generate/resource"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/spf13/cobra"
)

type Command struct {
	cmd     *cobra.Command
	Manager component.Manager
}

func New(cmd *cobra.Command, config config.Configuration) *Command {
	return &Command{
		cmd:     cmd,
		Manager: component.NewManager(config),
	}
}

func (c *Command) Setup() {
	generateCmd := &cobra.Command{
		Use:   "gen",
		Short: "Generates components",
	}

	brickCmd := &cobra.Command{
		Use:   "bck",
		Short: "Generates pure components (or bricks)",
	}

	brickCmd.AddCommand(brick.RunModel(c.Manager))
	brickCmd.AddCommand(brick.RunService(c.Manager))
	brickCmd.AddCommand(brick.RunHandler(c.Manager))
	brickCmd.AddCommand(brick.RunRepository(c.Manager))

	resourceCmd := &cobra.Command{
		Use:   "rsc",
		Short: "Generates conventional resources, grouping bricks",
	}

	resourceCmd.AddCommand(resource.RunRest(c.Manager))
	resourceCmd.AddCommand(resource.RunService(c.Manager))
	resourceCmd.AddCommand(resource.RunCore(c.Manager))

	generateCmd.AddCommand(brickCmd)
	generateCmd.AddCommand(resourceCmd)

	c.cmd.AddCommand(generateCmd)
}
