package molecule

import (
	"fmt"

	"github.com/charmingruby/bob/internal/command/generate/atom"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/input"
	"github.com/charmingruby/bob/internal/command/shared/fs"
	"github.com/spf13/cobra"
)

func RunRest(m component.Manager) *cobra.Command {
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

	if err := fs.GenerateFile(makeRestRegistryAtomComponent(
		fmt.Sprintf("%s/%s", moduleDir, "transport/rest/endpoint"),
		fmt.Sprintf("%s/%s", projectData, destinationDirectory),
		module,
	)); err != nil {
		panic(err)
	}

	if err := fs.GenerateFile(atom.MakeHandlerComponent(
		destinationDirectory,
		module,
		"ping",
	)); err != nil {
		panic(err)
	}
}

type restRegistryAtomData struct {
	Module          string
	SourceDirectory string
}

func makeRestRegistryAtomComponent(destinationDirectory, sourceDirectory, module string) fs.File {
	return atom.MakeRegistryComponent(atom.RegistryParams{
		Module:       module,
		TemplateName: "rest_registry",
		TemplateData: restRegistryAtomData{
			Module:          module,
			SourceDirectory: sourceDirectory,
		},
		RegistryName:         "endpoint",
		DestinationDirectory: destinationDirectory,
	})
}
