package atom

import (
	"github.com/charmingruby/bob/internal/command/gen/atom/constant"
	"github.com/charmingruby/bob/internal/command/gen/atom/structure"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
	"github.com/charmingruby/bob/internal/command/shared/opt"
	"github.com/charmingruby/bob/internal/command/shared/scaffold"
)

func MakePersistenceRepositoryComponent(m filesystem.Manager, module, name, database string) filesystem.File {
	var repoDB string = database

	if !opt.IsDatabaseOption(database) {
		repoDB = "exampledb"
	}

	repoDir := repoDB + "_repository"

	prepareDirectoriesForPersistenceRepository(m, module, repoDir)

	return component.New(component.ComponentInput{
		DestinationDirectory: scaffold.PersistencePath(m.ModuleDirectory(module), []string{repoDir}),
		Package:              module,
		Name:                 name,
		Suffix:               "repository",
	}).Componetize(component.ComponetizeInput{
		TemplateName: constant.REPOSITORY_UNIMPLEMENTED_TEMPLATE,
		TemplateData: structure.NewUnimplementedRepositoryData(m.DependencyPath(), module, name, repoDB),
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
