package unit

import (
	"github.com/charmingruby/bob/internal/component/architecture/unit/data"
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeUnimplementedRepository(m filesystem.Manager, module, name, database string) filesystem.File {
	prepareDirectoriesForUnimplementedRepository(m, module, database)

	template := "architecture/unit/repository/unimplemented"

	return base.New(base.ComponentInput{
		DestinationDirectory: definition.PersistencePath(m.ModuleDirectory(module), []string{database}),
		Package:              module,
		Name:                 name,
		Suffix:               "repository",
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			TemplateData: data.NewUnimplementedRepositoryData(m.DependencyPath(), module, name, database),
			FileName:     name,
			FileSuffix:   "repository",
		})
}

func prepareDirectoriesForUnimplementedRepository(m filesystem.Manager, module, database string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{definition.PERSISTENCE_PACKAGE, database},
	)
}
