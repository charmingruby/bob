package atom

import (
	"github.com/charmingruby/bob/internal/component/atom/constant"
	"github.com/charmingruby/bob/internal/component/atom/data"
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeModel(m filesystem.Manager, module, name string) filesystem.File {
	prepareDirectoriesForModel(m, module)

	return base.New(base.ComponentInput{
		DestinationDirectory: definition.CorePath(m.ModuleDirectory(module), []string{definition.MODEL_PACKAGE}),
		Package:              module,
		Name:                 name,
		HasTest:              true,
	}).Componetize(
		definition.ADD_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.MODEL_TEMPLATE,
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
