package resource

import (
	"github.com/charmingruby/bob/internal/cli/command/gen/resource/postgres"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/spf13/cobra"
)

func SetupCMD(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "res",
		Short: "Generates resources",
	}

	cmd.AddCommand(
		postgres.SetupCMD(fs),
	)

	return cmd
}
