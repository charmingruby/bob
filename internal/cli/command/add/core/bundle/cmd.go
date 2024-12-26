package bundle

import (
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "bundle",
		Aliases: []string{"bd"},
		Short:   "Generates bundles (aliases: bd)",
		Long:    "This command generates various units, which are individual business rules and components.",
	}

	cmd.AddCommand(RunService(fs))
	cmd.AddCommand(RunCore(fs))

	return cmd
}
