package bundle

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/core/bundle/service"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunService(m filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "service",
		Aliases: []string{"svc"},
		Short:   "Generates a new service bundle (aliases: svc)",
		Long:    "This command generates a new service bundle, which includes business logic and the necessary constructors.",
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
					Name:     "RepositoryName",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("Repository name")},
					Validate: survey.Required,
				},
				{
					Name:     "ModelName",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("Model to be managed name")},
					Validate: survey.Required,
				},
			}

			answers := struct {
				Module         string
				ServiceName    string
				RepositoryName string
				ModelName      string
			}{}

			survey.Ask(questions, &answers)

			components, err := service.Perfom(m, answers.RepositoryName, answers.Module, answers.ServiceName, answers.ModelName)
			if err != nil {
				output.ShutdownWithError(err.Error())
			}

			for _, c := range components {
				output.ComponentCreated(c.Identifier)
			}

			output.CommandSuccess("service bundle")
		},
	}

	return cmd
}
