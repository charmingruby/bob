package initialize

import (
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/setup/configure"
	"github.com/spf13/cobra"
)

func New(cmd *cobra.Command) {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Generates configuration files",
		Long:  "This command generates the necessary configuration files for the project.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := configure.PerformConfigure(); err != nil {
				output.ShutdownWithError(err.Error())
			}

			output.CommandSuccess("init")
		},
	}

	cmd.AddCommand(initCmd)
}
