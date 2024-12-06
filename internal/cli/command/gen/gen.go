package gen

import (
	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/cli/command/gen/atom"
	"github.com/charmingruby/bob/internal/cli/command/gen/molecule"
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

	atomCmd := &cobra.Command{
		Use:   "atm",
		Short: "Generates pure components (or atoms)",
	}

	atomCmd.AddCommand(atom.RunModel(c.fs))
	atomCmd.AddCommand(atom.RunService(c.fs))
	atomCmd.AddCommand(atom.RunRepository(c.fs))

	moleculeCmd := &cobra.Command{
		Use:   "mol",
		Short: "Generates conventional molecules, grouping atoms",
	}

	moleculeCmd.AddCommand(molecule.RunRest(c.fs))
	moleculeCmd.AddCommand(molecule.RunHandler(c.fs))
	moleculeCmd.AddCommand(molecule.RunService(c.fs))
	moleculeCmd.AddCommand(molecule.RunCore(c.fs))

	genCmd.AddCommand(atomCmd)
	genCmd.AddCommand(moleculeCmd)

	c.cmd.AddCommand(genCmd)
}
