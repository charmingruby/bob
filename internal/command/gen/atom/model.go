package atom

import (
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/constant"
	"github.com/charmingruby/bob/internal/command/shared/component/input"
	"github.com/charmingruby/bob/internal/command/shared/fs"
	"github.com/spf13/cobra"
)

func RunModel(m component.Manager) *cobra.Command {
	var (
		module string
		name   string
	)

	cmd := &cobra.Command{
		Use:   "model",
		Short: "Generates a new model",
		Run: func(cmd *cobra.Command, args []string) {
			if err := input.ValidateDefaultCommandInput(module, name); err != nil {
				panic(err)
			}

			if err := fs.GenerateFile(MakeModelComponent(
				m,
				module,
				name,
			)); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&name, "name", "n", "", "model name")

	return cmd
}

func MakeModelComponent(m component.Manager, module, name string) fs.File {
	component := *New(ComponentInput{
		DestinationDirectory: m.AppendToModuleDirectory(module, ModelPath()),
		Module:               module,
		Name:                 name,
		Suffix:               "",
		HasTest:              true,
	}, WithDefaultTemplate())

	return NewFileFromAtom(component, constant.MODEL_TEMPLATE)
}

func ModelPath() string {
	return "core/model"
}
