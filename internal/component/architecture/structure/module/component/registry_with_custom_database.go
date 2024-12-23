package component

import (
	"github.com/charmingruby/bob/internal/component/architecture/structure/module/data"
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeRegistryWithCustomDatabase(m filesystem.Manager, module, repositoryModel, database string) filesystem.File {
	prepareDirectoriesForRegistryWithCustomDatabase(m, module)

	template := "architecture/structure/module/module_with_custom_database"

	path := m.SourceDirectory + "/" + module

	return base.New(base.ComponentInput{
		Package:              module,
		DestinationDirectory: path,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			TemplateData: data.NewModuleWithDatabaseData(
				m.DependencyPath(),
				module,
				database,
				repositoryModel,
			),
			FileName: module,
		})
}

func prepareDirectoriesForRegistryWithCustomDatabase(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{module},
	)
}
