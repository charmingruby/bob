package brick

import (
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/constant"
	"github.com/charmingruby/bob/internal/command/shared/fs"
	"github.com/charmingruby/bob/internal/command/shared/validator/input"
	"github.com/spf13/cobra"
)

func RunService(m component.Manager) *cobra.Command {
	var (
		module string
		name   string
	)

	cmd := &cobra.Command{
		Use:   "service",
		Short: "Generates a new service",
		Run: func(cmd *cobra.Command, args []string) {
			if err := input.ValidateDefaultCommandInput(module, name); err != nil {
				panic(err)
			}

			if err := fs.GenerateFile(MakeServiceComponent(
				m.SourceDirectory,
				module,
				name,
			)); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&name, "name", "n", "", "service name")

	return cmd
}

func MakeServiceComponent(sourceDirectory, module, name string) fs.File {
	component := New(ComponentInput{
		Module:    module,
		Name:      name,
		Suffix:    "service",
		Directory: component.ModulePath(sourceDirectory, module, ServicePath()),
		HasTest:   false,
	}, WithDefaultTemplate())

	return NewFileFromBrick(*component, constant.SERVICE_TEMPLATE)
}

func ServicePath() string {
	return "core/service"
}
