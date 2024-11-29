package gen

import (
	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/command/gen/atom"
	"github.com/charmingruby/bob/internal/command/gen/molecule"
	"github.com/charmingruby/bob/internal/command/gen/molecule/rest"
	"github.com/charmingruby/bob/internal/command/gen/molecule/service"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/spf13/cobra"
)

type Command struct {
	cmd     *cobra.Command
	Manager component.Manager
}

func New(cmd *cobra.Command, config config.Configuration) *Command {
	return &Command{
		cmd:     cmd,
		Manager: component.NewManager(config),
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

	atomCmd.AddCommand(atom.RunModel(c.Manager))
	atomCmd.AddCommand(atom.RunService(c.Manager))
	atomCmd.AddCommand(atom.RunHandler(c.Manager))
	atomCmd.AddCommand(atom.RunRepository(c.Manager))

	moleculeCmd := &cobra.Command{
		Use:   "mol",
		Short: "Generates conventional molecules, grouping atoms",
	}

	moleculeCmd.AddCommand(rest.RunRest(c.Manager))
	moleculeCmd.AddCommand(service.RunService(c.Manager))
	moleculeCmd.AddCommand(molecule.RunCore(c.Manager))

	genCmd.AddCommand(atomCmd)
	genCmd.AddCommand(moleculeCmd)

	c.cmd.AddCommand(genCmd)
}
