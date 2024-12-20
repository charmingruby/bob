package unit

import (
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/spf13/cobra"
)

func SetupCMD(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unit",
		Short: "Generates pure components (or atoms)",
	}

	cmd.AddCommand(
		RunModel(fs),
		RunService(fs),
		RunRepository(fs),
		RunUnimplementedRepository(fs),
	)

	return cmd
}
