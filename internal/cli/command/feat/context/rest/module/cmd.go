package module

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "module",
		Aliases: []string{"mod"},
		Short:   "Generates module sets (aliases: mod)",
		Long:    "This command generates various REST API modules, which are sets of fully functional components.",
		Run: func(cmd *cobra.Command, args []string) {
			section := "structure"

			baseModuleName := "New base module"
			customDBModuleName := "New module with custom DB"
			postgresDBModuleName := "New module with Postgres DB"

			var templateChoice string
			prompt := &survey.Select{
				Message: input.ChooseSectionMessage(section),
				Options: []string{baseModuleName, customDBModuleName, postgresDBModuleName},
			}

			survey.AskOne(prompt, &templateChoice)

			switch templateChoice {
			case baseModuleName:
				RunBase(fs).Execute()
			case postgresDBModuleName:
				RunCustomDB(fs).Execute()
			case customDBModuleName:
				RunPostgresDB(fs).Execute()
			}
		},
	}

	return cmd
}
