package molecule

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/molecule/rest"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/spf13/cobra"
)

func RunRest(m filesystem.Manager) *cobra.Command {
	var (
		module string
	)

	cmd := &cobra.Command{
		Use:   "rest",
		Short: "Generates a new rest molecule",
		Run: func(cmd *cobra.Command, args []string) {
			if err := input.ValidateOnlyModuleCommandInput(module); err != nil {
				panic(err)
			}

			rest.MakeRest(m, module)
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")

	return cmd
}
