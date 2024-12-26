package rest

import (
	"github.com/charmingruby/bob/internal/cli/command/add/rest/bundle"
	"github.com/charmingruby/bob/internal/cli/command/add/rest/module"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "rest",
		Aliases: []string{"rs"},
		Short:   "Generates REST components (aliases: rs)",
		Long:    "This command generates REST components, which are sets of base o REST APIs."}

	cmd.AddCommand(
		bundle.SetupCmd(fs),
		module.SetupCmd(fs),
	)

	return cmd
}
