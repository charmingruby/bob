package generate

import (
	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/command/generate/brick"
	"github.com/charmingruby/bob/internal/command/generate/resource"
	"github.com/spf13/cobra"
)

type Command struct {
	cmd     *cobra.Command
	Project Project
}
type Project struct {
	Data            string
	SourceDirectory string
}

func New(cmd *cobra.Command, config config.Configuration) *Command {
	return &Command{
		cmd: cmd,
		Project: Project{
			Data:            config.BaseConfiguration.BaseURL + "/" + config.BaseConfiguration.ProjectName,
			SourceDirectory: config.BaseConfiguration.RootDir + "/" + config.BaseConfiguration.SourceDir,
		},
	}
}

func (c *Command) Setup() {
	generateCmd := &cobra.Command{
		Use:   "gen",
		Short: "Generates components",
	}

	brickCmd := &cobra.Command{
		Use:   "brick",
		Short: "Generates pure components",
	}

	brickCmd.AddCommand(brick.RunModel(c.Project.SourceDirectory))
	brickCmd.AddCommand(brick.RunService(c.Project.SourceDirectory))
	brickCmd.AddCommand(brick.RunHandler(c.Project.SourceDirectory))
	brickCmd.AddCommand(brick.RunRepository(c.Project.Data, c.Project.SourceDirectory))

	resourceCmd := &cobra.Command{
		Use:   "resource",
		Short: "Generates conventional services, grouping bricks",
	}

	resourceCmd.AddCommand(resource.RunRest(c.Project.Data, c.Project.SourceDirectory))
	resourceCmd.AddCommand(resource.RunService(c.Project.SourceDirectory))
	resourceCmd.AddCommand(resource.RunCore(c.Project.Data, c.Project.SourceDirectory))

	generateCmd.AddCommand(brickCmd)
	generateCmd.AddCommand(resourceCmd)

	c.cmd.AddCommand(generateCmd)
}

func (c *Command) ModuleDirectory(module string) string {
	return c.Project.SourceDirectory + "/" + module
}
