package unit

import (
	"github.com/charmingruby/bob/internal/component/architecture/unit/constant"
	"github.com/charmingruby/bob/internal/component/architecture/unit/data"
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeService(m filesystem.Manager, module, name string) filesystem.File {
	prepareDirectoriesForService(m, module)

	return base.New(base.ComponentInput{
		Package:              module,
		Name:                 name,
		Suffix:               "service",
		DestinationDirectory: definition.CorePath(m.ModuleDirectory(module), []string{definition.SERVICE_PACKAGE}),
	}).Componetize(
		definition.ADD_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.SERVICE_TEMPLATE,
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
