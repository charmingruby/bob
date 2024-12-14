package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/organism/module/constant"
	"github.com/charmingruby/bob/internal/component/organism/module/data"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/scaffold"
)

func MakeBaseRegistry(m filesystem.Manager, module string) filesystem.File {
	prepareDirectoriesForBaseRegistry(m, module)

	path := m.SourceDirectory + "/" + module

	return base.New(base.ComponentInput{
		Package:              module,
		DestinationDirectory: path,
	}).Componetize(
		scaffold.GENERATE_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.BASE_MODULE_TEMPLATE,
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
