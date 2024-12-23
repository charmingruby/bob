package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/resource/database/postgres/data"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakePostgresRepository(m filesystem.Manager, module, model string) filesystem.File {

	prepareDirectoriesForRepository(m, module, definition.POSTGRES_PACKAGE)

	template := "resource/database/postgres/repository"

	return base.New(base.ComponentInput{
		Package:              definition.POSTGRES_PACKAGE,
		DestinationDirectory: definition.PersistencePath(m.ModuleDirectory(module), []string{definition.POSTGRES_PACKAGE}),
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			TemplateData: data.NewPostgresRepositoryData(m.DependencyPath(), module, model),
			FileName:     model,
			FileSuffix:   "repository",
		})
}

func prepareDirectoriesForRepository(m filesystem.Manager, module, pkg string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{definition.PERSISTENCE_PACKAGE, pkg},
	)
}
