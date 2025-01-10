package resource

import (
	"github.com/charmingruby/bob/internal/cli/command/feat/resource/postgres"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "resource",
		Aliases: []string{"rsc"},
		Short:   "Generates resources (aliases: rsc)",
		Long:    "This command provides various infrastructure resource generation capabilities.",
	}

	cmd.AddCommand(
		postgres.SetupCmd(fs),
	)

	return cmd
}
