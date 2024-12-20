package postgres

import (
	"github.com/charmingruby/bob/internal/component/resource"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/spf13/cobra"
)

func RunDependecies(m filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deps",
		Short: "Generates a dependencies",
		Run: func(cmd *cobra.Command, args []string) {
			resource.PerformPostgresDependencies(m)
		},
	}

	return cmd
}
