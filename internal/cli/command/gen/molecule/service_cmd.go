package molecule

import (
	"fmt"

	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/molecule"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/spf13/cobra"
)

func RunService(m filesystem.Manager) *cobra.Command {
	var (
		module string
		repo   string
	)

	cmd := &cobra.Command{
		Use:   "service",
		Short: "Generates a new service molecule",
		Run: func(cmd *cobra.Command, args []string) {
			if err := m.GenerateNestedDirectories(
				fmt.Sprintf("%s/%s", m.SourceDirectory, module),
				[]string{"core", "service"},
			); err != nil {
				panic(err)
			}

			if err := input.ValidateOnlyModuleCommandInput(module); err != nil {
				panic(err)
			}

			molecule.MakeService(m, repo, module)
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&repo, "repo", "r", "", "repository dependency")

	return cmd
}
