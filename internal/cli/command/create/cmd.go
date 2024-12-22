package create

import (
	"github.com/charmingruby/bob/config"
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
	c.cmd.AddCommand(
		c.RunCreate(),
	)
}
