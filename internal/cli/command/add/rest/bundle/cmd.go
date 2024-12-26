package bundle

import (
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "bundle",
		Aliases: []string{"bd"},
		Short:   "Generates module sets (aliases: bd)",
		Long:    "This command generates various bundles, which are simpler sets of components.",
	}

	cmd.AddCommand(
		RunSetup(fs),
		RunHandler(fs),
	)

	return cmd
}
