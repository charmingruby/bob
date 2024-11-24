package brick

import (
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/constant"
	"github.com/charmingruby/bob/internal/command/shared/fs"
	"github.com/charmingruby/bob/internal/command/shared/validator/input"
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
				m.SourceDirectory,
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

func MakeModelComponent(sourceDirectory, module, name string) fs.File {
	component := *New(ComponentInput{
		Directory: component.ModulePath(sourceDirectory, module, ModelPath()),
		Module:    module,
		Name:      name,
		Suffix:    "",
		HasTest:   true,
	}, WithDefaultTemplate())

	return NewFileFromBrick(component, constant.MODEL_TEMPLATE)
}

func ModelPath() string {
	return "core/model"
}
