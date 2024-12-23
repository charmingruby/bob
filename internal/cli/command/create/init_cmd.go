package create

import (
	"github.com/charmingruby/bob/internal/component/setup"
	"github.com/spf13/cobra"
)

func Init(cmd *cobra.Command) {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Creates a new project",
		Run: func(cmd *cobra.Command, args []string) {
			setup.PerformConfigure()
		},
	}

	cmd.AddCommand(initCmd)
}
