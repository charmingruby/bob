package unit

import (
	"github.com/charmingruby/bob/internal/component/architecture/unit/data"
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeRepository(m filesystem.Manager, module, name string) filesystem.File {
	prepareDirectoriesForRepository(m, module)

	template := "architecture/unit/repository/contract"

	return base.New(base.ComponentInput{
		DestinationDirectory: definition.CorePath(m.ModuleDirectory(module), []string{definition.REPOSITORY_PACKAGE}),
		Package:              module,
		Name:                 name,
		Suffix:               "repository",
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			TemplateData: data.NewDependentPackageData(m.DependencyPath(), module, name),
			FileName:     name,
			FileSuffix:   "repository",
		})
}

func prepareDirectoriesForRepository(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{definition.CORE_PACKAGE, definition.REPOSITORY_PACKAGE},
	)
}
