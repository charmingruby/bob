package component

import (
	"github.com/charmingruby/bob/internal/component/architecture/bundle/rest/constant"
	"github.com/charmingruby/bob/internal/component/architecture/bundle/rest/data"
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeHandler(m filesystem.Manager, module, name string) filesystem.File {
	prepareDirectoriesForHandler(m, module)

	return base.New(base.ComponentInput{
		Package: module,
		Name:    name,
		Suffix:  "handler",
		DestinationDirectory: definition.TransportPath(
			m.ModuleDirectory(module),
			definition.REST_PACKAGE,
			[]string{definition.HANDLER_PACKAGE},
		),
	}).Componetize(
		definition.ADD_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.REST_HANDLER_TEMPLATE,
			TemplateData: data.NewHandlerData(
				m.DependencyPath(),
				module,
				name,
			),
			FileName:   name,
			FileSuffix: "handler",
		})
}

func prepareDirectoriesForHandler(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{definition.TRANSPORT_PACKAGE, definition.REST_PACKAGE, definition.HANDLER_PACKAGE},
	)
}
