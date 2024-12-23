package template

import (
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCMD(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "template",
		Short: "Postgres resources",
	}

	cmd.AddCommand(
		RunBase(fs),
		RunPostgres(fs),
	)

	return cmd
}
