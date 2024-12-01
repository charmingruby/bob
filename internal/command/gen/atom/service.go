package atom

import (
	"github.com/charmingruby/bob/internal/command/gen/atom/structure"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/constant"
	"github.com/charmingruby/bob/internal/command/shared/component/input"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
	"github.com/spf13/cobra"
)

func ServicePath() string {
	return "core/service"
}

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

			if err := filesystem.GenerateFile(MakeServiceComponent(
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

func MakeServiceComponent(sourceDirectory, module, name string) filesystem.File {
	return component.New(component.ComponentInput{
		Module:               module,
		Name:                 name,
		Suffix:               "service",
		DestinationDirectory: component.ModulePath(sourceDirectory, module, ServicePath()),
	}).Componetize(component.ComponetizeInput{
		TemplateName: constant.SERVICE_TEMPLATE,
		TemplateData: structure.NewDefaultData(name),
		FileName:     name,
		FileSuffix:   "service",
	})
}
