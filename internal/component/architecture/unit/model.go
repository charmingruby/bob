package unit

import (
	"github.com/charmingruby/bob/internal/component/architecture/unit/data"
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeModel(m filesystem.Manager, module, name string) filesystem.File {
	prepareDirectoriesForModel(m, module)

	template := "architecture/unit/model"

	return base.New(base.ComponentInput{
		DestinationDirectory: definition.CorePath(m.ModuleDirectory(module), []string{definition.MODEL_PACKAGE}),
		Package:              module,
		Name:                 name,
		HasTest:              true,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			TemplateData: data.NewDefaultData(name),
			FileName:     name,
		})
}

func prepareDirectoriesForModel(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{definition.CORE_PACKAGE, definition.MODEL_PACKAGE},
	)
}
