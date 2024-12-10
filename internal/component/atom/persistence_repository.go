package atom

import (
	"github.com/charmingruby/bob/internal/component/atom/constant"
	"github.com/charmingruby/bob/internal/component/atom/data"
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/scaffold"
)

func MakePersistenceRepository(m filesystem.Manager, module, name, database string) filesystem.File {
	repoDir := database + "_repository"

	prepareDirectoriesForPersistenceRepository(m, module, repoDir)

	return makeUnimplementedRepository(m, module, name, repoDir, database)
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
