package template

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/setup"
	"github.com/charmingruby/bob/internal/shared/filesystem"

	"github.com/spf13/cobra"
)

func RunPostgres(m filesystem.Manager) *cobra.Command {
	var goVersion string

	cmd := &cobra.Command{
		Use:   "pg",
		Short: "Creates a new project with PostgreSQL",
		Long:  "This command creates a new project using a PostgreSQL template.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := parseCreateInput(goVersion); err != nil {
				panic(err)
			}

			setup.PerformPostgresTemplate(m, goVersion)
		},
	}

	cmd.Flags().StringVarP(&goVersion, "golang version", "v", "1.23.3", "golang version for setup")

	return cmd
}

func parseCreateInput(goVersion string) error {
	args := []input.Arg{
		{
			FieldName: "go version",
			Value:     goVersion,
		},
	}

	return input.Validate(args)
}
