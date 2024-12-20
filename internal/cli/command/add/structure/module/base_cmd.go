package module

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/organism"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunModule(m filesystem.Manager) *cobra.Command {
	var (
		module    string
		modelName string
	)

	cmd := &cobra.Command{
		Use:   "base",
		Short: "Generates a base module",
		Run: func(cmd *cobra.Command, args []string) {
			if err := input.ValidateDefaultCommandInput(module, modelName); err != nil {
				panic(err)
			}

			organism.PerformBaseModule(m, module, modelName)
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&modelName, "model", "n", "", "model name")

	return cmd
}
