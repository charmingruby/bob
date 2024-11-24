package brick

import (
	"fmt"

	"github.com/charmingruby/bob/internal/command/shared/constant"
	"github.com/charmingruby/bob/internal/command/shared/fs"
	"github.com/charmingruby/bob/internal/command/shared/validator/input"
	"github.com/spf13/cobra"
)

func RunRepository(projectData, destinationDirectory string) *cobra.Command {
	var (
		module string
		name   string
	)

	cmd := &cobra.Command{
		Use:   "repository",
		Short: "Generates a new repository",
		Run: func(cmd *cobra.Command, args []string) {
			if err := input.ValidateDefaultCommandInput(module, name); err != nil {
				panic(err)
			}

			if err := fs.GenerateFile(MakeRepositoryComponent(
				fmt.Sprintf("%s/%s", projectData, destinationDirectory),
				destinationDirectory,
				module,
				name,
			)); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&name, "name", "n", "", "model to be managed by the repository")

	return cmd
}

func MakeRepositoryComponent(sourcePath, destinationDirectory, module, name string) fs.File {
	component := *New(ComponentInput{
		Directory: fmt.Sprintf("%s/%s/core/%s",
			destinationDirectory,
			module,
			constant.REPOSITORY_TEMPLATE,
		),
		Module:  module,
		Name:    name,
		Suffix:  "repository",
		HasTest: false,
	}, WithModuleDependenciesTemplate(sourcePath))

	return NewFileFromBrick(component, constant.REPOSITORY_TEMPLATE)
}
