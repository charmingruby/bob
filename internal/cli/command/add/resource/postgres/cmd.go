package postgres

import (
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "postgres",
		Aliases: []string{"pg"},
		Short:   "Postgres resources (aliases: pg)",
		Long:    "This command provides various PostgreSQL resources."}

	cmd.AddCommand(
		RunRepo(fs),
		RunDeps(fs),
		RunMig(fs),
	)

	return cmd
}
