package resource

import (
	"fmt"

	"github.com/charmingruby/bob/internal/command/generate/brick"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/fs"
	"github.com/charmingruby/bob/internal/command/shared/validator/input"
	"github.com/spf13/cobra"
)

func RunRest(m component.Manager) *cobra.Command {
	var (
		module string
	)

	cmd := &cobra.Command{
		Use:   "rest",
		Short: "Generates a new rest resource",
		Run: func(cmd *cobra.Command, args []string) {
			if err := input.ValidateOnlyModuleCommandInput(module); err != nil {
				panic(err)
			}

			MakeRest(m.Data, m.SourceDirectory, module)
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")

	return cmd
}

func MakeRest(projectData, destinationDirectory, module string) {
	moduleDir := fmt.Sprintf("%s/%s", destinationDirectory, module)

	if err := fs.GenerateNestedDirectories(
		moduleDir,
		[]string{"transport", "rest", "endpoint"},
	); err != nil {
		panic(err)
	}

	if err := fs.GenerateFile(makeRestRegistryBrickComponent(
		fmt.Sprintf("%s/%s", moduleDir, "transport/rest/endpoint"),
		fmt.Sprintf("%s/%s", projectData, destinationDirectory),
		module,
	)); err != nil {
		panic(err)
	}

	if err := fs.GenerateFile(brick.MakeHandlerComponent(
		destinationDirectory,
		module,
		"ping",
	)); err != nil {
		panic(err)
	}
}

type restRegistryBrickData struct {
	Module          string
	SourceDirectory string
}

func makeRestRegistryBrickComponent(destinationDirectory, sourceDirectory, module string) fs.File {
	return makeRegistryBrick(registryBrickParams{
		Module:       module,
		TemplateName: "rest_registry",
		TemplateData: restRegistryBrickData{
			Module:          module,
			SourceDirectory: sourceDirectory,
		},
		RegistryName:         "endpoint",
		DestinationDirectory: destinationDirectory,
	})
}
