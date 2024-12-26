package rest

import (
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "rest",
		Aliases: []string{"rs"},
		Short:   "Generates REST API project from templates (aliases: rs)",
		Long:    "This command provides various template resources to generate a REST API project.",
	}

	cmd.AddCommand(
		RunBase(fs),
		RunPostgres(fs),
	)

	return cmd
}
