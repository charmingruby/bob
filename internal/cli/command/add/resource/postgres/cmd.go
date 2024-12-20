package postgres

import (
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCMD(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pg",
		Short: "Postgres resources",
	}

	cmd.AddCommand(
		RunRepository(fs),
		RunDependecies(fs),
		RunMigration(fs),
	)

	return cmd
}
