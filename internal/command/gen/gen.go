package gen

import (
	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/command/gen/atom"
	"github.com/charmingruby/bob/internal/command/gen/molecule"
	"github.com/charmingruby/bob/internal/command/gen/molecule/rest"
	"github.com/charmingruby/bob/internal/command/gen/molecule/service"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
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

	atomCmd := &cobra.Command{
		Use:   "atm",
		Short: "Generates pure components (or atoms)",
	}

	atomCmd.AddCommand(atom.RunModel(c.fs))
	atomCmd.AddCommand(atom.RunService(c.fs))
	atomCmd.AddCommand(atom.RunHandler(c.fs))
	atomCmd.AddCommand(atom.RunRepository(c.fs))

	moleculeCmd := &cobra.Command{
		Use:   "mol",
		Short: "Generates conventional molecules, grouping atoms",
	}

	moleculeCmd.AddCommand(rest.RunRest(c.fs))
	moleculeCmd.AddCommand(service.RunService(c.fs))
	moleculeCmd.AddCommand(molecule.RunCore(c.fs))

	genCmd.AddCommand(atomCmd)
	genCmd.AddCommand(moleculeCmd)

	c.cmd.AddCommand(genCmd)
}
