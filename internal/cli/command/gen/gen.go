package gen

import (
	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/cli/command/gen/atom"
	"github.com/charmingruby/bob/internal/cli/command/gen/molecule"
	"github.com/charmingruby/bob/internal/cli/command/gen/organism"
	"github.com/charmingruby/bob/internal/cli/command/gen/resource"
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
	genCmd := &cobra.Command{
		Use:   "gen",
		Short: "Generates components",
	}

	genCmd.AddCommand(
		atom.SetupCMD(c.fs),
		molecule.SetupCMD(c.fs),
		resource.SetupCMD(c.fs),
		organism.SetupCMD(c.fs),
	)

	c.cmd.AddCommand(genCmd)
}
