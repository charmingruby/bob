package bundle

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/core/bundle/core"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunCore(m filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "core",
		Aliases: []string{"cr"},
		Short:   "Generates a new core bundle (aliases: cr)",
		Long:    "This command generates a new core bundle, which includes domain rules, models, contracts, and other core components.",
		Run: func(cmd *cobra.Command, args []string) {
			questions := []*survey.Question{
				{
					Name:     "Module",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("Module")},
					Validate: survey.Required,
				},
				{
					Name:     "ModelName",
					Prompt:   &survey.Input{Message: input.EnterValueMessage("Base model name")},
					Validate: survey.Required,
				},
			}

			answers := struct {
				Module    string
				ModelName string
			}{}

			survey.Ask(questions, &answers)

			components, err := core.Perform(m, answers.Module, answers.ModelName)
			if err != nil {
				output.ShutdownWithError(err.Error())
			}

			for _, c := range components {
				output.ComponentCreated(c.Identifier)
			}

			output.CommandSuccess("core bundle")
		},
	}

	return cmd
}
