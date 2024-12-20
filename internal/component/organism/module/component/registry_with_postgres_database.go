package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/organism/module/constant"
	"github.com/charmingruby/bob/internal/component/organism/module/data"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeRegistryWithPostgresDatabase(m filesystem.Manager, module, repositoryModel string) filesystem.File {
	prepareDirectoriesForRegistryWithPostgresDatabase(m, module)

	path := m.SourceDirectory + "/" + module

	return base.New(base.ComponentInput{
		Package:              module,
		DestinationDirectory: path,
	}).Componetize(
		definition.ADD_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.MODULE_WITH_POSTGRES_DATABASE_TEMPLATE,
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
