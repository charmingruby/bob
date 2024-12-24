package add

import (
	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/cli/command/add/bundle"
	"github.com/charmingruby/bob/internal/cli/command/add/resource"
	"github.com/charmingruby/bob/internal/cli/command/add/structure/module"
	"github.com/charmingruby/bob/internal/cli/command/add/unit"
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
	}

	cmd.AddCommand(
		unit.SetupCmd(c.fs),
		bundle.SetupCmd(c.fs),
		module.SetupCmd(c.fs),
		resource.SetupCmd(c.fs),
	)

	c.cmd.AddCommand(cmd)
}
