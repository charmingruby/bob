package organism

import (
	"github.com/charmingruby/bob/internal/cli/command/add/structure/module"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "structure",
		Short: "Generates organisms",
	}

	cmd.AddCommand(
		module.SetupCmd(fs),
	)

	return cmd
}
