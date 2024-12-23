package template

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/setup"
	"github.com/charmingruby/bob/internal/shared/filesystem"

	"github.com/spf13/cobra"
)

func RunBase(m filesystem.Manager) *cobra.Command {
	var goVersion string

	cmd := &cobra.Command{
		Use:   "base",
		Short: "Creates a new project",
		Run: func(cmd *cobra.Command, args []string) {
			if err := parseBaseInput(goVersion); err != nil {
				panic(err)
			}

			setup.PerfomBaseTemplate(m, goVersion, "dynamo")
		},
	}

	cmd.Flags().StringVarP(&goVersion, "golang version", "v", "1.23.3", "golang version for setup, by default, it will be 1.23.3")

	return cmd
}

func parseBaseInput(goVersion string) error {
	args := []input.Arg{
		{
			FieldName: "go version",
			Value:     goVersion,
		},
	}

	return input.Validate(args)
}
