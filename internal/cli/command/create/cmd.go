package create

import (
	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/cli/command/create/template"
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
		Use:   "create",
		Short: "Creates a new project",
		Long:  "This command creates a new project by setting up the project structure from templates.",
	}

	cmd.AddCommand(
		template.SetupCmd(c.fs),
	)

	c.cmd.AddCommand(cmd)
}
