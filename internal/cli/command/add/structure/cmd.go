package organism

import (
	"github.com/charmingruby/bob/internal/cli/command/add/structure/module"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/spf13/cobra"
)

func SetupCMD(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "structure",
		Short: "Generates organisms",
	}

	cmd.AddCommand(
		module.SetupCMD(fs),
	)

	return cmd
}
