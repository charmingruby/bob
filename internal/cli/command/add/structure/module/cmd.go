package module

import (
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/spf13/cobra"
)

func SetupCMD(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "module",
		Short: "Module generator",
	}

	cmd.AddCommand(
		RunModule(fs),
		RunModuleWithPostgresDatabase(fs),
		RunModuleWithCustomDatabase(fs),
	)

	return cmd
}
