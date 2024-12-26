package unit

import (
	"fmt"

	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

type unimplementedRepositoryData struct {
	SourcePath string
	Module     string
	Name       string
	Database   string
}

func newUnimplementedRepositoryData(sourcePath, module, name, db string) unimplementedRepositoryData {
	return unimplementedRepositoryData{
		SourcePath: sourcePath,
		Module:     base.ModuleFormat(module),
		Name:       base.PublicNameFormat(name),
		Database:   base.ModuleFormat(db),
	}
}

func MakeUnimplementedRepository(m filesystem.Manager, module, name, database string) filesystem.File {
	prepareDirectoriesForUnimplementedRepository(m, module, database)

	template := TemplatePath("repository/unimplemented")

	destination := definition.PersistencePath(m.ModuleDirectory(module), []string{database})

	content := fmt.Sprintf("%s unimplemented repository", name)

	return base.New(base.ComponentInput{
		Identifier:           base.BuildIdentifier(module, content, destination),
		Package:              module,
		Name:                 name,
		Suffix:               "repository",
		DestinationDirectory: destination,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			TemplateData: newUnimplementedRepositoryData(m.DependencyPath(), module, name, database),
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
