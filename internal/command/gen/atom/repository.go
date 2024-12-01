package atom

import (
	"github.com/charmingruby/bob/internal/command/gen/atom/structure"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/constant"
	"github.com/charmingruby/bob/internal/command/shared/component/input"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
	"github.com/spf13/cobra"
)

func RepositoryPath() string {
	return "core/repository"
}

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

			if err := filesystem.GenerateFile(MakeRepositoryComponent(
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

func MakeRepositoryComponent(m component.Manager, module, name string) filesystem.File {
	return component.New(component.ComponentInput{
		DestinationDirectory: component.ModulePath(m.SourceDirectory, module, RepositoryPath()),
		Module:               module,
		Name:                 name,
		Suffix:               "repository",
	}).Componetize(component.ComponetizeInput{
		TemplateName: constant.REPOSITORY_TEMPLATE,
		TemplateData: structure.NewDependentPackageData(m.DependencyPath(module), module, name),
		FileName:     name,
		FileSuffix:   "repository",
	})
}
