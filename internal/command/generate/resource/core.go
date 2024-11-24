package resource

import (
	"fmt"

	"github.com/charmingruby/bob/internal/command/generate/brick"
	"github.com/charmingruby/bob/internal/command/shared/fs"
	"github.com/charmingruby/bob/internal/command/shared/validator/input"
	"github.com/spf13/cobra"
)

func RunCore(projectData, destinationDirectory string) *cobra.Command {
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

			MakeCore(projectData, destinationDirectory, module)
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")

	return cmd
}

func MakeCore(projectData, destinationDirectory, module string) {
	moduleDir := fmt.Sprintf("%s/%s", destinationDirectory, module)
	sourcePath := fmt.Sprintf("%s/%s", projectData, destinationDirectory)

	if err := fs.GenerateDirectory(moduleDir, "core"); err != nil {
		panic(err)
	}

	if err := fs.GenerateMultipleDirectories(
		fmt.Sprintf("%s/core", moduleDir),
		[]string{"service", "model", "repository"},
	); err != nil {
		panic(err)
	}

	MakeService(destinationDirectory, module)

	println(destinationDirectory)

	if err := fs.GenerateFile(brick.MakeRepositoryComponent(sourcePath, destinationDirectory, module, "example")); err != nil {
		panic(err)
	}

	if err := fs.GenerateFile(brick.MakeModelComponent(destinationDirectory, module, "example")); err != nil {
		panic(err)
	}
}
