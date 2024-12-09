package atom

import (
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/spf13/cobra"
)

func SetupCMD(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "atm",
		Short: "Generates pure components (or atoms)",
	}

	cmd.AddCommand(RunModel(fs))
	cmd.AddCommand(RunService(fs))
	cmd.AddCommand(RunRepository(fs))
	cmd.AddCommand(RunHandler(fs))

	return cmd
}
