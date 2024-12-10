package postgres

import (
	"github.com/charmingruby/bob/internal/component/resource/database/postgres/component"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/spf13/cobra"
)

func RunDependecies(m filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dependencies",
		Short: "Generates a dependencies",
		Run: func(cmd *cobra.Command, args []string) {
			conn := component.MakePostgresConnection(m)
			if err := m.GenerateFile(conn); err != nil {
				panic(err)
			}
		},
	}

	return cmd
}
