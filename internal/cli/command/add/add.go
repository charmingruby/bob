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
	genCmd := &cobra.Command{
		Use:   "add",
		Short: "Generates components",
	}

	genCmd.AddCommand(
		unit.SetupCMD(c.fs),
		bundle.SetupCMD(c.fs),
		module.SetupCMD(c.fs),
		resource.SetupCMD(c.fs),
	)

	c.cmd.AddCommand(genCmd)
}
