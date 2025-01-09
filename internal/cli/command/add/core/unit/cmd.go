package unit

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "unit",
		Aliases: []string{"ut"},
		Short:   "Generates unit (aliases: ut)",
		Long:    "This command generates the smallest and core components fo architecture.",
		Run: func(cmd *cobra.Command, args []string) {
			section := "structure"

			modelName := "Model"
			serviceName := "Service"
			repoName := "Repository contract"
			unimplRepo := "Unimplemented repository"

			var templateChoice string
			prompt := &survey.Select{
				Message: input.ChooseSectionMessage(section),
				Options: []string{modelName, serviceName, repoName, unimplRepo},
			}
			survey.AskOne(prompt, &templateChoice)

			switch templateChoice {
			case modelName:
				RunModel(fs).Execute()
			case serviceName:
				RunService(fs).Execute()
			case repoName:
				RunRepo(fs).Execute()
			case unimplRepo:
				RunUnimplRepo(fs).Execute()
			}
		},
	}

	return cmd
}
