package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/molecule/rest/constant"
	"github.com/charmingruby/bob/internal/component/molecule/rest/data"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/shared"
)

func MakeHandler(m filesystem.Manager, module, name string) filesystem.File {
	prepareDirectoriesForHandler(m, module)

	return base.New(base.ComponentInput{
		Package: module,
		Name:    name,
		Suffix:  "handler",
		DestinationDirectory: shared.TransportPath(
			m.ModuleDirectory(module),
			shared.REST_PACKAGE,
			[]string{shared.HANDLER_PACKAGE},
		),
	}).Componetize(
		shared.GENERATE_COMMAND,
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
		[]string{shared.TRANSPORT_PACKAGE, shared.REST_PACKAGE, shared.HANDLER_PACKAGE},
	)
}
