package component

import (
	"github.com/charmingruby/bob/internal/component/architecture/structure/module/data"
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeRegistryWithPostgresDatabase(m filesystem.Manager, module, repositoryModel string) filesystem.File {
	prepareDirectoriesForRegistryWithPostgresDatabase(m, module)

	template := "architecture/structure/module/module_with_postgres_database"

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
				"postgres",
				repositoryModel,
			),
			FileName: module,
		})
}

func prepareDirectoriesForRegistryWithPostgresDatabase(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{module},
	)
}
