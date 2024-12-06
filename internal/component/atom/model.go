package atom

import (
	"github.com/charmingruby/bob/internal/component/atom/data"
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/scaffold"
)

func MakeModelComponent(m filesystem.Manager, module, name string) filesystem.File {
	prepareDirectoriesForModel(m, module)

	return base.New(base.ComponentInput{
		DestinationDirectory: scaffold.CorePath(m.ModuleDirectory(module), []string{scaffold.MODEL_PACKAGE}),
		Package:              module,
		Name:                 name,
		HasTest:              true,
	}).Componetize(base.ComponetizeInput{
		TemplateName: MODEL_TEMPLATE,
		TemplateData: data.NewDefaultData(name),
		FileName:     name,
	})
}

func prepareDirectoriesForModel(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{scaffold.CORE_PACKAGE, scaffold.MODEL_PACKAGE},
	)
}
