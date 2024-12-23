package component

import (
	"github.com/charmingruby/bob/internal/component/architecture/structure/module/data"
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeBaseRegistry(m filesystem.Manager, module string) filesystem.File {
	prepareDirectoriesForBaseRegistry(m, module)

	template := "architecture/structure/module/base_module"

	path := m.SourceDirectory + "/" + module

	return base.New(base.ComponentInput{
		Package:              module,
		DestinationDirectory: path,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			TemplateData: data.NewBaseModuleData(
				m.DependencyPath(),
				module,
			),
			FileName: module,
		})
}

func prepareDirectoriesForBaseRegistry(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{module},
	)
}
