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
		module               string
		serviceName          string
		repoName             string
		modelToBeManagedName string
	)

	cmd := &cobra.Command{
		Use:     "service",
		Aliases: []string{"svc"},
		Short:   "Generates a new service bundle (aliases: svc)",
		Long:    "This command generates a new service bundle, which includes business logic and the necessary constructors.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := parseServiceInput(module, repoName, modelToBeManagedName, serviceName); err != nil {
				output.ShutdownWithError(err.Error())
			}

			components, err := service.Perfom(m, repoName, module, serviceName, modelToBeManagedName)
			if err != nil {
				output.ShutdownWithError(err.Error())
			}

			for _, c := range components {
				output.ComponentCreated(c.Identifier)
			}

			output.CommandSuccess("service bundle")
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module")
	cmd.Flags().StringVarP(&repoName, "repo", "r", "", "repository name dependency")
	cmd.Flags().StringVarP(&repoName, "model name", "e", "", "model to be managed by service")
	cmd.Flags().StringVarP(&repoName, "service name", "n", "", "service name")

	return cmd
}

func parseServiceInput(module, repo, modelToBeManagedName, serviceName string) error {
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
		{
			FieldName:  "service name",
			Value:      serviceName,
			IsRequired: true,
			Type:       input.StringType,
		},
		{
			FieldName:  "model name",
			Value:      modelToBeManagedName,
			IsRequired: true,
			Type:       input.StringType,
		},
	}

	return input.Validate(inputs)
}
