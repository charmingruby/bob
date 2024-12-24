package module

import (
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "module",
		Aliases: []string{"mod"},
		Short:   "Generates modules (aliases: mod)",
		Long:    "This command generates modules, which are sets of actions separated by bounded contexts.",
	}

	cmd.AddCommand(
		RunBase(fs),
		RunPostgresDB(fs),
		RunCustomDB(fs),
	)

	return cmd
}
