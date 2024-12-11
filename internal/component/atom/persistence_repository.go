package atom

import (
	"github.com/charmingruby/bob/internal/component/atom/constant"
	"github.com/charmingruby/bob/internal/component/atom/data"
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/scaffold"
)

func MakePersistenceRepository(m filesystem.Manager, module, name, database string) filesystem.File {
	prepareDirectoriesForPersistenceRepository(m, module, database)

	return makeUnimplementedRepository(m, module, name, database)
}

func makeUnimplementedRepository(m filesystem.Manager, module, name, database string) filesystem.File {
	return base.New(base.ComponentInput{
		DestinationDirectory: scaffold.PersistencePath(m.ModuleDirectory(module), []string{database}),
		Package:              module,
		Name:                 name,
		Suffix:               "repository",
	}).Componetize(base.ComponetizeInput{
		TemplateName: constant.REPOSITORY_UNIMPLEMENTED_TEMPLATE,
		TemplateData: data.NewUnimplementedRepositoryData(m.DependencyPath(), module, name, database),
		FileName:     name,
		FileSuffix:   "repository",
	})
}

func prepareDirectoriesForPersistenceRepository(m filesystem.Manager, module, database string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{scaffold.PERSISTENCE_PACKAGE, database},
	)
}
