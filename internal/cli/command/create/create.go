package create

import (
	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/organism"
	"github.com/charmingruby/bob/internal/filesystem"

	"github.com/spf13/cobra"
)

type Command struct {
	cmd *cobra.Command
	fs  filesystem.Manager
}

func New(cmd *cobra.Command, config config.Configuration) *Command {
	return &Command{
		cmd: cmd,
		fs:  filesystem.New(config),
	}
}

func (c *Command) Setup() {
	var goVersion string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a new project",
		Run: func(cmd *cobra.Command, args []string) {
			if err := ValidateNewCommandInput(goVersion); err != nil {
				panic(err)
			}

			organism.MakeAndRunSetup(c.fs, goVersion)
		},
	}

	cmd.Flags().StringVarP(&goVersion, "golang version", "v", "1.23.3", "golang version for setup, by default, it will be 1.23.3")

	c.cmd.AddCommand(cmd)
}

func ValidateNewCommandInput(goVersion string) error {
	args := []input.Arg{
		{
			FieldName: "go version",
			Value:     goVersion,
		},
	}

	if err := input.ValidateArgsList(args); err != nil {
		return err
	}

	return nil
}
