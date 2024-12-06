package rest_component

import (
	"github.com/charmingruby/bob/internal/command/gen/molecule/rest/constant"
	"github.com/charmingruby/bob/internal/command/gen/molecule/rest/structure"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
	"github.com/charmingruby/bob/internal/command/shared/scaffold"
)

func MakeHandlerComponent(m filesystem.Manager, module, name string) filesystem.File {
	prepareDirectoriesForHandler(m, module)

	return component.New(component.ComponentInput{
		Package: module,
		Name:    name,
		Suffix:  "handler",
		DestinationDirectory: scaffold.TransportPath(
			m.ModuleDirectory(module),
			scaffold.REST_PACKAGE,
			[]string{scaffold.HANDLER_PACKAGE},
		),
	}).Componetize(component.ComponetizeInput{
		TemplateName: constant.REST_HANDLER_TEMPLATE,
		TemplateData: structure.NewHandlerData(
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
