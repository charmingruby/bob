package rest_component

import (
	"github.com/charmingruby/bob/internal/command/gen/molecule/rest/constant"
	"github.com/charmingruby/bob/internal/command/gen/molecule/rest/structure"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
	"github.com/charmingruby/bob/internal/command/shared/scaffold"
)

func MakeRequestHelperComponent(m filesystem.Manager) filesystem.File {
	prepareDirectoriesForRequestHelper(m, scaffold.SHARED_MODULE)

	return component.New(component.ComponentInput{
		Package: scaffold.REST_PACKAGE,
		DestinationDirectory: scaffold.TransportPath(
			m.ModuleDirectory(scaffold.SHARED_MODULE),
			scaffold.REST_PACKAGE,
			nil,
		),
	}).Componetize(component.ComponetizeInput{
		TemplateName: constant.REST_REQUEST_HELPER_TEMPLATE,
		TemplateData: structure.NewRequestHelperData(m.DependencyPath()),
		FileName:     "request",
	})
}

func prepareDirectoriesForRequestHelper(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{module, scaffold.TRANSPORT_PACKAGE, scaffold.REST_PACKAGE},
	)
}
