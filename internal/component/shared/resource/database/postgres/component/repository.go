package component

import (
	"fmt"

	"github.com/charmingruby/bob/internal/component/shared/resource/database/postgres/data"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakePostgresRepository(m filesystem.Manager, module, model string) filesystem.File {
	prepareDirectoriesForRepository(m, module, definition.POSTGRES_PACKAGE)

	template := "resource/database/postgres/repository"

	destination := definition.PersistencePath(m.ModuleDirectory(module), []string{definition.POSTGRES_PACKAGE})

	content := fmt.Sprintf("%s %s repository", definition.POSTGRES_PACKAGE, model)

	return base.New(base.ComponentInput{
		Identifier:           base.BuildIdentifier(module, content, destination),
		Package:              definition.POSTGRES_PACKAGE,
		DestinationDirectory: destination,
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
