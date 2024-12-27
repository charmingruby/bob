package component

import (
	"fmt"

	"github.com/charmingruby/bob/internal/component/context/rest"
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
		Module:     base.SnakeCaseFormat(module),
	}
}

func MakeRegistry(m filesystem.Manager, module string) filesystem.File {
	prepareDirectoriesForRegistry(m, module)

	template := rest.TemplatePath("module/base")

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

func prepareDirectoriesForRegistry(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{module},
	)
}
