package module

import (
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "module",
		Short: "Module generator",
	}

	cmd.AddCommand(
		RunBase(fs),
		RunPostgresDB(fs),
		RunCustomDB(fs),
	)

	return cmd
}
