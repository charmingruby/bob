package component

import (
	"fmt"

	"github.com/charmingruby/bob/internal/component/context/rest/module/shared/data"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeRegistryWithCustomDatabase(m filesystem.Manager, module, repositoryModel, database string) filesystem.File {
	prepareDirectoriesForRegistryWithCustomDatabase(m, module)

	template := "architecture/structure/module/module_with_custom_database"

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
