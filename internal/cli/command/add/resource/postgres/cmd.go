package postgres

import (
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pg",
		Short: "Postgres resources",
		Long:  "This command provides various PostgreSQL resources."}

	cmd.AddCommand(
		RunRepo(fs),
		RunDeps(fs),
		RunMig(fs),
	)

	return cmd
}
