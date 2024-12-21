package bundle

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/molecule"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunService(m filesystem.Manager) *cobra.Command {
	var (
		module   string
		repoName string
	)

	cmd := &cobra.Command{
		Use:   "service",
		Short: "Generates a new service molecule",
		Run: func(cmd *cobra.Command, args []string) {
			if err := parseServiceInput(module, repoName); err != nil {
				panic(err)
			}

			molecule.PerformService(m, repoName, module)
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&repoName, "repo", "r", "", "repository name dependency")

	return cmd
}

func parseServiceInput(module, repo string) error {
	inputs := []input.Arg{
		{
			FieldName:  "module",
			IsRequired: true,
			Value:      module,
			Type:       input.StringType,
		},
		{
			FieldName: "repo",
			Value:     repo,
			Type:      input.StringType,
		},
	}

	return input.Validate(inputs)
}
