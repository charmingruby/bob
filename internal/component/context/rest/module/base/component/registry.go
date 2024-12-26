package component

import (
	"fmt"

	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

type baseModuleData struct {
	SourcePath string
	Module     string
}

func newBaseModuleData(sourcePath, module string) baseModuleData {
	return baseModuleData{
		SourcePath: sourcePath,
		Module:     base.ModuleFormat(module),
	}
}

func MakeBaseRegistry(m filesystem.Manager, module string) filesystem.File {
	prepareDirectoriesForBaseRegistry(m, module)

	template := "architecture/structure/module/base_module"

	destination := m.SourceDirectory + "/" + module

	content := fmt.Sprintf("%s module", module)

	return base.New(base.ComponentInput{
		Identifier:           base.BuildIdentifier(module, content, destination),
		Package:              module,
		DestinationDirectory: destination,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			TemplateData: newBaseModuleData(
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
