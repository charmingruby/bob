package atom

import (
	"github.com/charmingruby/bob/internal/component/atom/constant"
	"github.com/charmingruby/bob/internal/component/atom/data"
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/library"
	"github.com/charmingruby/bob/internal/component/shared/opt"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/scaffold"
)

func MakePersistenceRepository(m filesystem.Manager, module, name, database string) filesystem.File {
	var repoDB string = database

	if !opt.IsDatabaseOption(database) {
		repoDB = "exampledb"
	}

	repoDir := repoDB + "_repository"

	prepareDirectoriesForPersistenceRepository(m, module, repoDir)

	switch database {
	case opt.POSTGRES_DATABASE:
		library.MakePostgresDependencies(m)
		return library.MakePostgresRepository(m, module, name)
	default:
		return makeUnimplementedRepository(m, module, name, repoDir, repoDB)
	}
}

func makeUnimplementedRepository(m filesystem.Manager, module, name, repoDir, repoDB string) filesystem.File {
	return base.New(base.ComponentInput{
		DestinationDirectory: scaffold.PersistencePath(m.ModuleDirectory(module), []string{repoDir}),
		Package:              module,
		Name:                 name,
		Suffix:               "repository",
	}).Componetize(base.ComponetizeInput{
		TemplateName: constant.REPOSITORY_UNIMPLEMENTED_TEMPLATE,
		TemplateData: data.NewUnimplementedRepositoryData(m.DependencyPath(), module, name, repoDB),
		FileName:     name,
		FileSuffix:   "repository",
	})
}

func prepareDirectoriesForPersistenceRepository(m filesystem.Manager, module, repoDir string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{scaffold.PERSISTENCE_PACKAGE, repoDir},
	)
}
