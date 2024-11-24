package resource

import (
	"fmt"

	"github.com/charmingruby/bob/internal/command/generate/brick"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/fs"
	"github.com/charmingruby/bob/internal/command/shared/validator/input"
	"github.com/spf13/cobra"
)

func RunService(m component.Manager) *cobra.Command {
	var (
		module string
	)

	cmd := &cobra.Command{
		Use:   "service",
		Short: "Generates a new service resource",
		Run: func(cmd *cobra.Command, args []string) {
			if err := fs.GenerateNestedDirectories(
				fmt.Sprintf("%s/%s", m.SourceDirectory, module),
				[]string{"core", "service"},
			); err != nil {
				panic(err)
			}

			if err := input.ValidateOnlyModuleCommandInput(module); err != nil {
				panic(err)
			}

			MakeService(m.SourceDirectory, module)
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")

	return cmd
}

func MakeService(sourceDirectory, module string) {
	registryDirectory := sourceDirectory + "/" + module + "/core/service"

	if err := fs.GenerateFile(makeServiceRegistryBrickComponent(
		registryDirectory,
		module)); err != nil {
		panic(err)
	}

	if err := fs.GenerateFile(brick.MakeServiceComponent(sourceDirectory, module, "example")); err != nil {
		panic(err)
	}
}

type serviceRegistryBrickData struct {
	Module string
}

func makeServiceRegistryBrickComponent(destinationDirectory, module string) fs.File {
	return makeRegistryBrick(registryBrickParams{
		Module:       module,
		TemplateName: "service_registry",
		TemplateData: serviceRegistryBrickData{
			Module: module,
		},
		RegistryName:         "service",
		DestinationDirectory: destinationDirectory,
	})
}

func ServicePath() string {
	return "core/service"
}
