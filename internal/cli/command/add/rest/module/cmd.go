package module

import (
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "module",
		Aliases: []string{"mod"},
		Short:   "Generates module sets (aliases: mod)",
		Long:    "This command generates various REST API modules, which are sets of fully functional components.",
	}

	cmd.AddCommand(
		RunBase(fs),
		RunCustomDB(fs),
		RunPostgresDB(fs),
	)

	return cmd
}
