package module

import (
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "module",
		Short: "Generates modules",
		Long:  "This command generates modules, which are sets of actions separated by bounded contexts.",
	}

	cmd.AddCommand(
		RunBase(fs),
		RunPostgresDB(fs),
		RunCustomDB(fs),
	)

	return cmd
}
