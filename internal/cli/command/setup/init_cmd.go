package setup

import (
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/config/schema"
	"github.com/spf13/cobra"
)

func RunInit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Generates configuration files",
		Long:  "This command generates the necessary configuration files for the project.",
		Run: func(cmd *cobra.Command, args []string) {
			component, err := schema.Perform()
			if err != nil {
				output.ShutdownWithError(err.Error())
			}

			output.ComponentCreated(component.Identifier)

			output.CommandSuccess("init")
		},
	}

	return cmd
}
