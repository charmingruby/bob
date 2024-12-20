package bundle

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/molecule"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/spf13/cobra"
)

func RunCore(m filesystem.Manager) *cobra.Command {
	var (
		module    string
		modelName string
	)

	cmd := &cobra.Command{
		Use:   "core",
		Short: "Generates a new core molecule",
		Run: func(cmd *cobra.Command, args []string) {
			if err := input.ValidateDefaultCommandInput(module, modelName); err != nil {
				panic(err)
			}

			molecule.MakeAndRunCore(m, module, modelName)
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&modelName, "model", "n", "", "model name")

	return cmd
}
