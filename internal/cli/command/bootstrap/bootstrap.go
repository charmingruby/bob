package bootstrap

import (
	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/component/organism"
	"github.com/charmingruby/bob/internal/filesystem"

	"github.com/spf13/cobra"
)

type Command struct {
	cmd *cobra.Command
	fs  filesystem.Manager
}

func New(cmd *cobra.Command, config config.Configuration) *Command {
	return &Command{
		cmd: cmd,
		fs:  filesystem.New(config),
	}
}

func (c *Command) Setup() {
	cmd := &cobra.Command{
		Use:   "new",
		Short: "Creates a new project",
		Run: func(cmd *cobra.Command, args []string) {
			organism.MakeAndRunSetup(c.fs)
		},
	}

	c.cmd.AddCommand(cmd)
}
