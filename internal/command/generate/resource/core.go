package resource

import (
	"fmt"

	"github.com/charmingruby/bob/internal/command/generate/brick"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/fs"
	"github.com/charmingruby/bob/internal/command/shared/validator/input"
	"github.com/spf13/cobra"
)

func RunCore(m component.Manager) *cobra.Command {
	var (
		module string
	)

	cmd := &cobra.Command{
		Use:   "core",
		Short: "Generates a new core resource",
		Run: func(cmd *cobra.Command, args []string) {
			if err := input.ValidateOnlyModuleCommandInput(module); err != nil {
				panic(err)
			}

			MakeCore(m.Data, m.SourceDirectory, module)
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")

	return cmd
}

func MakeCore(projectData, sourceDirectory, module string) {
	moduleDir := fmt.Sprintf("%s/%s", sourceDirectory, module)
	sourcePath := fmt.Sprintf("%s/%s", projectData, sourceDirectory)

	if err := fs.GenerateDirectory(moduleDir, "core"); err != nil {
		panic(err)
	}

	if err := fs.GenerateMultipleDirectories(
		component.ModulePath(sourceDirectory, module, CorePath()),
		[]string{"service", "model", "repository"},
	); err != nil {
		panic(err)
	}

	MakeService(sourceDirectory, module)

	if err := fs.GenerateFile(brick.MakeRepositoryComponent(sourceDirectory, module, "example", sourcePath)); err != nil {
		panic(err)
	}

	if err := fs.GenerateFile(brick.MakeModelComponent(sourceDirectory, module, "example")); err != nil {
		panic(err)
	}
}

func CorePath() string {
	return "core"
}
