package unit

import (
	"fmt"

	"github.com/charmingruby/bob/internal/component/architecture/unit/data"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeModel(m filesystem.Manager, module, name string) filesystem.File {
	prepareDirectoriesForModel(m, module)

	template := "architecture/unit/model"

	destination := definition.CorePath(m.ModuleDirectory(module), []string{definition.MODEL_PACKAGE})

	content := fmt.Sprintf("%s model", name)

	return base.New(base.ComponentInput{
		Identifier:           base.BuildIdentifier(module, content, destination),
		Package:              module,
		Name:                 name,
		DestinationDirectory: destination,
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
