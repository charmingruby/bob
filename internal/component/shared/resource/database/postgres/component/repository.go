package component

import (
	"fmt"

	"github.com/charmingruby/bob/internal/component/shared/resource/database"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

type postgresRepositoryData struct {
	Module         string
	SourcePath     string
	LowerCaseModel string
	UpperCaseModel string
}

func newPostgresRepositoryData(sourcePath, module, model string) postgresRepositoryData {
	return postgresRepositoryData{
		Module:         base.SnakeCaseFormat(module),
		SourcePath:     sourcePath,
		LowerCaseModel: base.LowerCaseFormat(model),
		UpperCaseModel: base.CapitalizedFormat(model),
	}
}

func MakePostgresRepository(m filesystem.Manager, module, model string) filesystem.File {
	prepareDirectoriesForRepository(m, module, definition.POSTGRES_PACKAGE)

	template := database.TemplatePath("postgres/repository")

	destination := definition.PersistencePath(m.ModuleDirectory(module), []string{definition.POSTGRES_PACKAGE})

	content := fmt.Sprintf("%s %s repository", definition.POSTGRES_PACKAGE, model)

	return base.New(base.ComponentInput{
		Identifier:           base.BuildIdentifier(module, content, destination),
		Package:              definition.POSTGRES_PACKAGE,
		DestinationDirectory: destination,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			TemplateData: newPostgresRepositoryData(m.DependencyPath(), module, model),
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
