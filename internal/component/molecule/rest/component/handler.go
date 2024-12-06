package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/molecule/constant"
	"github.com/charmingruby/bob/internal/component/molecule/rest/data"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/scaffold"
)

func MakeHandler(m filesystem.Manager, module, name string) filesystem.File {
	prepareDirectoriesForHandler(m, module)

	return base.New(base.ComponentInput{
		Package: module,
		Name:    name,
		Suffix:  "handler",
		DestinationDirectory: scaffold.TransportPath(
			m.ModuleDirectory(module),
			scaffold.REST_PACKAGE,
			[]string{scaffold.HANDLER_PACKAGE},
		),
	}).Componetize(base.ComponetizeInput{
		TemplateName: constant.REST_HANDLER_TEMPLATE,
		TemplateData: data.NewHandlerData(
			m.DependencyPath(),
			scaffold.SHARED_MODULE,
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
		[]string{scaffold.TRANSPORT_PACKAGE, scaffold.REST_PACKAGE, scaffold.HANDLER_PACKAGE},
	)
}
