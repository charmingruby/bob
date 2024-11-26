package atom

import (
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/constant"
	"github.com/charmingruby/bob/internal/command/shared/component/input"
	"github.com/charmingruby/bob/internal/command/shared/fs"
	"github.com/spf13/cobra"
)

func RunRepository(m component.Manager) *cobra.Command {
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
				m,
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

func MakeRepositoryComponent(m component.Manager, module, name string) fs.File {
	component := *New(ComponentInput{
		Directory: component.ModulePath(m.SourceDirectory, module, RepositoryPath()),
		Module:    module,
		Name:      name,
		Suffix:    "repository",
		HasTest:   false,
	}, WithModuleDependenciesTemplate(m.DependencyPath(module)))

	return NewFileFromAtom(component, constant.REPOSITORY_TEMPLATE)
}

func RepositoryPath() string {
	return "core/repository"
}
