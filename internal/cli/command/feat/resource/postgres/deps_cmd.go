package postgres

import (
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/shared/resource/database/postgres"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunDeps(m filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deps",
		Short: "Generates PostgreSQL dependencies",
		Long:  "This command generates all necessary dependencies for PostgreSQL",
		Run: func(cmd *cobra.Command, args []string) {
			components, err := postgres.PerformDependencies(m)
			if err != nil {
				output.ShutdownWithError(err.Error())
			}

			for _, c := range components {
				output.ComponentCreated(c.Identifier)
			}

			output.CommandSuccess("postgres dependencies")
		},
	}

	return cmd
}
