package unit

import (
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "unit",
		Aliases: []string{"ut"},
		Short:   "Generates unit (aliases: ut)",
		Long:    "This command generates the smallest and core components fo architecture.",
	}

	cmd.AddCommand(
		RunModel(fs),
		RunService(fs),
		RunRepo(fs),
		RunUnimplRepo(fs),
	)

	return cmd
}
