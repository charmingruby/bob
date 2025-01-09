package unit

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/core/unit"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunService(m filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "service",
		Aliases: []string{"svc"},
		Short:   "Generates a new service (aliases: svc)",
		Long:    "This command generates a new service to build business logic.",
		Run: func(cmd *cobra.Command, args []string) {
			questions := []*survey.Question{
				{
					Name:     "Module",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("Module")},
					Validate: survey.Required,
				},
				{
					Name:     "ServiceName",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("Service name")},
					Validate: survey.Required,
				},
				{
					Name:     "ModelName",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("Model to be managed name")},
					Validate: survey.Required,
				},
			}

			answers := struct {
				Module      string
				ServiceName string
				ModelName   string
			}{}

			survey.Ask(questions, &answers)

			service := unit.MakeService(m, answers.Module, answers.ServiceName, answers.ModelName)

			if err := m.GenerateFile(service); err != nil {
				output.ShutdownWithError(err.Error())
			}

			output.ComponentCreated(service.Identifier)
			output.CommandSuccess("service unit")
		},
	}

	return cmd
}
