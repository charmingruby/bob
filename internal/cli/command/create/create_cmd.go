package create

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/setup"

	"github.com/spf13/cobra"
)

func (c *Command) RunCreate() *cobra.Command {
	var goVersion string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a new project",
		Run: func(cmd *cobra.Command, args []string) {
			if err := parseCreateInput(goVersion); err != nil {
				panic(err)
			}

			setup.PerformScaffold(c.fs, goVersion)
		},
	}

	cmd.Flags().StringVarP(&goVersion, "golang version", "v", "1.23.3", "golang version for setup, by default, it will be 1.23.3")

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
