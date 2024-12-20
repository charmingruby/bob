package bundle

import (
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCMD(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bundle",
		Short: "Generates conventional molecules, grouping atoms",
	}

	cmd.AddCommand(RunRest(fs))
	cmd.AddCommand(RunService(fs))
	cmd.AddCommand(RunCore(fs))

	return cmd
}
