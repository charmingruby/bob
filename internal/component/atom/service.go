package atom

import (
	"github.com/charmingruby/bob/internal/component/atom/data"
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/scaffold"
)

func MakeServiceComponent(m filesystem.Manager, module, name string) filesystem.File {
	prepareDirectoriesForService(m, module)

	return base.New(base.ComponentInput{
		Package:              module,
		Name:                 name,
		Suffix:               "service",
		DestinationDirectory: scaffold.CorePath(m.ModuleDirectory(module), []string{scaffold.SERVICE_PACKAGE}),
	}).Componetize(base.ComponetizeInput{
		TemplateName: SERVICE_TEMPLATE,
		TemplateData: data.NewDefaultData(name),
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
