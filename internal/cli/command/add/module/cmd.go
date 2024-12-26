package module

import (
	module "github.com/charmingruby/bob/internal/cli/command/add/module/rest"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "module",
		Aliases: []string{"mod"},
		Short:   "Generates module sets (aliases: mod)",
		Long:    "This command generates various module sets.",
	}

	cmd.AddCommand(
		module.SetupCmd(fs),
	)

	return cmd
}
