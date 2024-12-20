package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/resource/database/postgres/constant"
	"github.com/charmingruby/bob/internal/component/resource/database/postgres/data"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakePostgresRepository(m filesystem.Manager, module, model string) filesystem.File {
	prepareDirectoriesForRepository(m, module)

	return base.New(base.ComponentInput{
		Package:              constant.POSTGRES_PACKAGE,
		DestinationDirectory: definition.PersistencePath(m.ModuleDirectory(module), []string{constant.POSTGRES_PACKAGE}),
	}).Componetize(
		definition.ADD_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.POSTGRES_REPOSITORY_TEMPLATE,
			TemplateData: data.NewPostgresRepositoryData(m.DependencyPath(), module, model),
			FileName:     model,
			FileSuffix:   "repository",
		})
}

func prepareDirectoriesForRepository(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{definition.PERSISTENCE_PACKAGE, constant.POSTGRES_PACKAGE},
	)
}
