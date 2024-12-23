package unit

import (
	"fmt"

	"github.com/charmingruby/bob/internal/component/architecture/unit/data"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeService(m filesystem.Manager, module, name string) filesystem.File {
	prepareDirectoriesForService(m, module)

	template := "architecture/unit/service"

	destination := definition.CorePath(m.ModuleDirectory(module), []string{definition.SERVICE_PACKAGE})

	content := fmt.Sprintf("%s service", name)

	return base.New(base.ComponentInput{
		Identifier:           base.BuildIdentifier(module, content, destination),
		Package:              module,
		Name:                 name,
		Suffix:               "service",
		DestinationDirectory: destination,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			TemplateData: data.NewDefaultData(name),
			FileName:     name,
			FileSuffix:   "service",
		})
}

func prepareDirectoriesForService(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{definition.CORE_PACKAGE, definition.SERVICE_PACKAGE},
	)
}
