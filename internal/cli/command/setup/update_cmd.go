package setup

import (
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/setup/update"
	"github.com/spf13/cobra"
)

func RunUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update Bob version",
		Long:  "This command updates the Bob version by downloading and executing the latest installation script.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := update.Perform(); err != nil {
				output.ShutdownWithError(err.Error())
			}

			output.CommandSuccess("update")
		},
	}

	return cmd
}
