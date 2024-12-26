package core

import (
	"github.com/charmingruby/bob/internal/cli/command/add/core/bundle"
	"github.com/charmingruby/bob/internal/cli/command/add/core/unit"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "core",
		Aliases: []string{"cr"},
		Short:   "Generates core components (aliases: cr)",
		Long:    "This command generates various module sets where the business rules are defined.",
	}

	cmd.AddCommand(
		bundle.SetupCmd(fs),
		unit.SetupCmd(fs),
	)

	return cmd
}
