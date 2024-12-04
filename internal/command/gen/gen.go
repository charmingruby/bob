package gen

import (
	"github.com/charmingruby/bob/config"
	atomCommand "github.com/charmingruby/bob/internal/command/gen/atom/cmd"
	coreCommand "github.com/charmingruby/bob/internal/command/gen/molecule/core/cmd"
	restCommand "github.com/charmingruby/bob/internal/command/gen/molecule/rest/cmd"
	serviceCommand "github.com/charmingruby/bob/internal/command/gen/molecule/service/cmd"

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

	atomCmd.AddCommand(atomCommand.RunModel(c.fs))
	atomCmd.AddCommand(atomCommand.RunService(c.fs))
	atomCmd.AddCommand(atomCommand.RunRepository(c.fs))

	moleculeCmd := &cobra.Command{
		Use:   "mol",
		Short: "Generates conventional molecules, grouping atoms",
	}

	moleculeCmd.AddCommand(restCommand.RunRest(c.fs))
	moleculeCmd.AddCommand(restCommand.RunHandler(c.fs))
	moleculeCmd.AddCommand(serviceCommand.RunService(c.fs))
	moleculeCmd.AddCommand(coreCommand.RunCore(c.fs))

	genCmd.AddCommand(atomCmd)
	genCmd.AddCommand(moleculeCmd)

	c.cmd.AddCommand(genCmd)
}
