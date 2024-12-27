package add

import (
	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/cli/command/add/core"
	"github.com/charmingruby/bob/internal/cli/command/add/resource"
	"github.com/charmingruby/bob/internal/cli/command/add/rest"
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
		core.SetupCmd(c.fs),
		resource.SetupCmd(c.fs),
		rest.SetupCmd(c.fs),
	)

	c.cmd.AddCommand(cmd)
}
