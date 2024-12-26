package module

import (
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "rest",
		Aliases: []string{"rs"},
		Short:   "Generates REST modules (aliases: rs)",
		Long:    "This command generates REST modules, which are sets of base o REST APIs."}

	cmd.AddCommand(
		RunBase(fs),
		RunPostgresDB(fs),
		RunCustomDB(fs),
	)

	return cmd
}
