package postgres

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func SetupCmd(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "postgres",
		Aliases: []string{"pg"},
		Short:   "Postgres resources (aliases: pg)",
		Long:    "This command provides various PostgreSQL resources.",
		Run: func(cmd *cobra.Command, args []string) {
			section := "template options"

			repoName := "New repository"
			migrationName := "New migration"
			dependencies := "Install dependencies"

			var templateChoice string
			prompt := &survey.Select{
				Message: input.ChooseSectionMessage(section),
				Options: []string{migrationName, repoName, dependencies},
			}
			survey.AskOne(prompt, &templateChoice)

			switch templateChoice {
			case migrationName:
				RunMig(fs).Execute()
			case repoName:
				RunRepo(fs).Execute()
			case dependencies:
				RunDeps(fs).Execute()
			}
		},
	}

	return cmd
}
