package generate

import (
	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/command/generate/brick"
	"github.com/spf13/cobra"
)

type Command struct {
	cmd    *cobra.Command
	config config.Configuration
}

func New(cmd *cobra.Command, config config.Configuration) *Command {
	return &Command{
		cmd:    cmd,
		config: config,
	}
}

func (c *Command) Setup() {
	destinationDirectory := c.config.BaseConfiguration.RootDir + "/" + c.config.BaseConfiguration.SourceDir

	generateCmd := &cobra.Command{
		Use:   "gen",
		Short: "Generates components",
	}

	generateCmd.AddCommand(brick.RunModel(destinationDirectory))
	generateCmd.AddCommand(brick.RunService(destinationDirectory))
	generateCmd.AddCommand(brick.RunHandler(destinationDirectory))

	resourceCmd := &cobra.Command{
		Use:   "resource",
		Short: "Generates resources",
	}

	//resourceCmd.AddCommand(resource.RunRest(c.config))

	generateCmd.AddCommand(resourceCmd)

	c.cmd.AddCommand(generateCmd)
}
