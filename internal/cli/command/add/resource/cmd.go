package resource

import (
	"github.com/charmingruby/bob/internal/cli/command/add/resource/postgres"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resource",
		Short: "Generates resources",
	}

	cmd.AddCommand(
		postgres.SetupCmd(fs),
	)

	return cmd
}
