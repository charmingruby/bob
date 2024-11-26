package atom

import (
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/constant"
	"github.com/charmingruby/bob/internal/command/shared/component/input"
	"github.com/charmingruby/bob/internal/command/shared/fs"
	"github.com/spf13/cobra"
)

func RunHandler(m component.Manager) *cobra.Command {
	var (
		module string
		name   string
	)

	cmd := &cobra.Command{
		Use:   "handler",
		Short: "Generates a new handler",
		Run: func(cmd *cobra.Command, args []string) {
			if err := input.ValidateDefaultCommandInput(module, name); err != nil {
				panic(err)
			}

			if err := fs.GenerateFile(MakeHandlerComponent(
				m.SourceDirectory,
				module,
				name,
			)); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&name, "name", "n", "", "handler name")

	return cmd
}

func MakeHandlerComponent(sourceDirectory, module, name string) fs.File {
	component := New(ComponentInput{
		Module:    module,
		Name:      name,
		Suffix:    "handler",
		Directory: component.ModulePath(sourceDirectory, module, HandlerPath()),
		HasTest:   false,
	}, WithDefaultTemplate())

	return NewFileFromAtom(*component, constant.HANDLER_TEMPLATE)
}

func HandlerPath() string {
	return "transport/rest/endpoint"
}
