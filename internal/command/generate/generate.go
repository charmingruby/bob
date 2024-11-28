package generate

import (
	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/command/generate/atom"
	"github.com/charmingruby/bob/internal/command/generate/molecule"
	"github.com/charmingruby/bob/internal/command/generate/molecule/rest"
	"github.com/charmingruby/bob/internal/command/generate/molecule/service"
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
	generateCmd := &cobra.Command{
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

	generateCmd.AddCommand(atomCmd)
	generateCmd.AddCommand(moleculeCmd)

	c.cmd.AddCommand(generateCmd)
}
