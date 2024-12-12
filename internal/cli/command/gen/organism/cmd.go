package organism

import (
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/spf13/cobra"
)

func SetupCMD(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "org",
		Short: "Generates organisms",
	}

	cmd.AddCommand(RunModule(fs))

	return cmd
}
