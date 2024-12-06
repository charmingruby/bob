package atom

import (
	"github.com/charmingruby/bob/internal/command/gen/atom/constant"
	"github.com/charmingruby/bob/internal/command/gen/atom/structure"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
	"github.com/charmingruby/bob/internal/command/shared/scaffold"
)

func MakeRepositoryComponent(m filesystem.Manager, module, name string) filesystem.File {
	prepareDirectoriesForRepository(m, module)

	return component.New(component.ComponentInput{
		DestinationDirectory: scaffold.CorePath(m.ModuleDirectory(module), []string{scaffold.REPOSITORY_PACKAGE}),
		Package:              module,
		Name:                 name,
		Suffix:               "repository",
	}).Componetize(component.ComponetizeInput{
		TemplateName: constant.REPOSITORY_TEMPLATE,
		TemplateData: structure.NewDependentPackageData(m.DependencyPath(), module, name),
		FileName:     name,
		FileSuffix:   "repository",
	})
}

func prepareDirectoriesForRepository(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{scaffold.CORE_PACKAGE, scaffold.REPOSITORY_PACKAGE},
	)
}
