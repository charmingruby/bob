package generate

import (
	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/command/generate/brick"
	"github.com/charmingruby/bob/internal/command/generate/resource"
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
	projectData := c.config.BaseConfiguration.BaseURL + "/" + c.config.BaseConfiguration.ProjectName
	destinationDirectory := c.config.BaseConfiguration.RootDir + "/" + c.config.BaseConfiguration.SourceDir

	generateCmd := &cobra.Command{
		Use:   "gen",
		Short: "Generates components",
	}

	brickCmd := &cobra.Command{
		Use:   "brick",
		Short: "Generates pure components",
	}

	brickCmd.AddCommand(brick.RunModel(destinationDirectory))
	brickCmd.AddCommand(brick.RunService(destinationDirectory))
	brickCmd.AddCommand(brick.RunHandler(destinationDirectory))

	resourceCmd := &cobra.Command{
		Use:   "resource",
		Short: "Generates conventional services, grouping bricks",
	}

	resourceCmd.AddCommand(resource.RunRest(projectData, destinationDirectory))
	resourceCmd.AddCommand(resource.RunService(destinationDirectory))

	generateCmd.AddCommand(brickCmd)
	generateCmd.AddCommand(resourceCmd)

	c.cmd.AddCommand(generateCmd)
}
