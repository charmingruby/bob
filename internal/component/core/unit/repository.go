package unit

import (
	"fmt"

	"github.com/charmingruby/bob/internal/component/core/unit/data"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeRepository(m filesystem.Manager, module, name string) filesystem.File {
	prepareDirectoriesForRepository(m, module)

	template := "architecture/unit/repository/contract"

	destination := definition.CorePath(m.ModuleDirectory(module), []string{definition.REPOSITORY_PACKAGE})

	content := fmt.Sprintf("%s repository", name)

	return base.New(base.ComponentInput{
		Identifier:           base.BuildIdentifier(module, content, destination),
		Package:              module,
		Name:                 name,
		DestinationDirectory: destination,
		Suffix:               "repository",
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			TemplateData: data.NewDependentPackageData(m.DependencyPath(), module, name),
			FileName:     name,
			FileSuffix:   "repository",
		})
}

func prepareDirectoriesForRepository(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{definition.CORE_PACKAGE, definition.REPOSITORY_PACKAGE},
	)
}