package atom

import (
	"github.com/charmingruby/bob/internal/command/gen/atom/constant"
	"github.com/charmingruby/bob/internal/command/gen/atom/structure"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
	"github.com/charmingruby/bob/internal/command/shared/scaffold"
)

func MakeServiceComponent(m filesystem.Manager, module, name string) filesystem.File {
	prepareDirectoriesForService(m, module)

	return component.New(component.ComponentInput{
		Package:              module,
		Name:                 name,
		Suffix:               "service",
		DestinationDirectory: scaffold.CorePath(m.ModuleDirectory(module), []string{scaffold.SERVICE_PACKAGE}),
	}).Componetize(component.ComponetizeInput{
		TemplateName: constant.SERVICE_TEMPLATE,
		TemplateData: structure.NewDefaultData(name),
		FileName:     name,
		FileSuffix:   "service",
	})
}

func prepareDirectoriesForService(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{scaffold.CORE_PACKAGE, scaffold.SERVICE_PACKAGE},
	)
}
