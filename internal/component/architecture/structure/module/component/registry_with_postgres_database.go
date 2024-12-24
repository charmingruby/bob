package component

import (
	"fmt"

	"github.com/charmingruby/bob/internal/component/architecture/structure/module/data"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeRegistryWithPostgresDatabase(m filesystem.Manager, module, repositoryModel string) filesystem.File {
	prepareDirectoriesForRegistryWithPostgresDatabase(m, module)

	template := "architecture/structure/module/module_with_postgres_database"

	destination := m.SourceDirectory + "/" + module

	content := fmt.Sprintf("%s module", module)

	return base.New(base.ComponentInput{
		Identifier:           base.BuildIdentifier(module, content, destination),
		Package:              module,
		DestinationDirectory: destination,
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
