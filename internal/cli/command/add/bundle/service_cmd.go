package bundle

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/core/bundle/service"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunService(m filesystem.Manager) *cobra.Command {
	var (
		module   string
		repoName string
	)

	cmd := &cobra.Command{
		Use:     "service",
		Aliases: []string{"svc"},
		Short:   "Generates a new service bundle (aliases: svc)",
		Long:    "This command generates a new service bundle, which includes business logic and the necessary constructors.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := parseServiceInput(module, repoName); err != nil {
				output.ShutdownWithError(err.Error())
			}

			if err := service.PerformService(m, repoName, module); err != nil {
				output.ShutdownWithError(err.Error())
			}

			output.CommandSuccess("service bundle")
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
