package postgres

import (
	"github.com/charmingruby/bob/internal/component/resource"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunDeps(m filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deps",
		Short: "Generates PostgreSQL dependencies",
		Long:  "This command generates all necessary dependencies for PostgreSQL",
		Run: func(cmd *cobra.Command, args []string) {
			resource.PerformPostgresDependencies(m)
		},
	}

	return cmd
}
