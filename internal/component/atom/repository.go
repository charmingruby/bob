package atom

import (
	"github.com/charmingruby/bob/internal/component/atom/constant"
	"github.com/charmingruby/bob/internal/component/atom/data"
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeRepository(m filesystem.Manager, module, name string) filesystem.File {
	prepareDirectoriesForRepository(m, module)

	return base.New(base.ComponentInput{
		DestinationDirectory: definition.CorePath(m.ModuleDirectory(module), []string{definition.REPOSITORY_PACKAGE}),
		Package:              module,
		Name:                 name,
		Suffix:               "repository",
	}).Componetize(
		definition.ADD_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.REPOSITORY_CONTRACT_TEMPLATE,
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
