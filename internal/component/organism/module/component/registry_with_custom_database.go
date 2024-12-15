package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/organism/module/constant"
	"github.com/charmingruby/bob/internal/component/organism/module/data"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/shared"
)

func MakeRegistryWithCustomDatabase(m filesystem.Manager, module, repositoryModel, database string) filesystem.File {
	prepareDirectoriesForRegistryWithCustomDatabase(m, module)

	path := m.SourceDirectory + "/" + module

	return base.New(base.ComponentInput{
		Package:              module,
		DestinationDirectory: path,
	}).Componetize(
		shared.GENERATE_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.MODULE_WITH_CUSTOM_DATABASE_TEMPLATE,
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
