package rest

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/context/rest/template/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"

	"github.com/spf13/cobra"
)

func RunBase(m filesystem.Manager) *cobra.Command {
	var goVersion string
	var database string

	cmd := &cobra.Command{
		Use:   "base",
		Short: "Creates a new project from a base template",
		Long:  "This command creates a new project using a customizable base persistence layer.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := parseBaseInput(goVersion, database); err != nil {
				output.ShutdownWithError(err.Error())
			}

			if err := base.PerfomBase(m, goVersion, database); err != nil {
				output.ShutdownWithError(err.Error())
			}

			output.CommandSuccess("base template")
		},
	}

	cmd.Flags().StringVarP(&goVersion, "golang version", "v", "1.23.3", "golang version for setup")
	cmd.Flags().StringVarP(&database, "database", "d", "", "base database to be implemented")

	return cmd
}

func parseBaseInput(goVersion, database string) error {
	args := []input.Arg{
		{
			FieldName: "go version",
			Value:     goVersion,
			Type:      input.StringType,
		},
		{
			FieldName:  "database",
			Value:      database,
			IsRequired: true,
			Type:       input.StringType,
		},
	}

	return input.Validate(args)
}
