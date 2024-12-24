package template

import (
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "template",
		Aliases: []string{"tpl"},
		Short:   "Generates project from templates (aliases: tmpl)",
		Long:    "This command provides various template resources.",
	}

	cmd.AddCommand(
		RunBase(fs),
		RunPostgres(fs),
	)

	return cmd
}
