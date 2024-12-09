package postgres

import (
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/spf13/cobra"
)

func SetupCMD(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pg",
		Short: "Postgres resources",
	}

	cmd.AddCommand(RunRepository(fs))

	return cmd
}
