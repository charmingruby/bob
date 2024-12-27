package unit

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/core/unit"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunService(m filesystem.Manager) *cobra.Command {
	var (
		module               string
		serviceName          string
		modelToBeManagedName string
	)

	cmd := &cobra.Command{
		Use:     "service",
		Aliases: []string{"svc"},
		Short:   "Generates a new service (aliases: svc)",
		Long:    "This command generates a new service to build business logic.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := parseServiceInput(module, serviceName, modelToBeManagedName); err != nil {
				output.ShutdownWithError(err.Error())
			}

			service := unit.MakeService(m, module, serviceName, modelToBeManagedName)

			if err := m.GenerateFile(service); err != nil {
				output.ShutdownWithError(err.Error())
			}

			output.ComponentCreated(service.Identifier)
			output.CommandSuccess("service unit")
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module")
	cmd.Flags().StringVarP(&serviceName, "service name", "n", "", "service name")
	cmd.Flags().StringVarP(&modelToBeManagedName, "model name", "e", "", "model to be managed name")

	return cmd
}

func parseServiceInput(module, serviceName, modelToBeManagedName string) error {
	inputs := []input.Arg{
		{
			FieldName:  "module",
			IsRequired: true,
			Value:      module,
			Type:       input.StringType,
		},
		{
			FieldName:  "service name",
			IsRequired: true,
			Value:      serviceName,
			Type:       input.StringType,
		},
		{
			FieldName:  "model name",
			IsRequired: true,
			Value:      modelToBeManagedName,
			Type:       input.StringType,
		},
	}

	return input.Validate(inputs)
}
