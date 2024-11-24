package resource

import (
	"fmt"

	"github.com/charmingruby/bob/internal/command/generate/brick"
	"github.com/charmingruby/bob/internal/command/shared/fs"
	"github.com/charmingruby/bob/internal/command/shared/validator/input"
	"github.com/spf13/cobra"
)

func RunRest(projectData, destinationDirectory string) *cobra.Command {
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

			MakeRest(projectData, destinationDirectory, module)
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
	Module     string
	SourcePath string
}

func makeRestRegistryBrickComponent(destinationDirectory, sourcePath, module string) fs.File {
	return makeRegistryBrick(registryBrickParams{
		Module:       module,
		TemplateName: "rest_registry",
		TemplateData: restRegistryBrickData{
			Module:     module,
			SourcePath: sourcePath,
		},
		RegistryName:         "endpoint",
		DestinationDirectory: destinationDirectory,
	})
}
