package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/molecule/constant"
	"github.com/charmingruby/bob/internal/component/molecule/rest/data"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/scaffold"
)

func MakeRequestHelper(m filesystem.Manager) filesystem.File {
	prepareDirectoriesForRequestHelper(m, scaffold.SHARED_MODULE)

	return base.New(base.ComponentInput{
		Package: scaffold.REST_PACKAGE,
		DestinationDirectory: scaffold.TransportPath(
			m.ModuleDirectory(scaffold.SHARED_MODULE),
			scaffold.REST_PACKAGE,
			nil,
		),
	}).Componetize(base.ComponetizeInput{
		TemplateName: constant.REST_REQUEST_HELPER_TEMPLATE,
		TemplateData: data.NewRequestHelperData(m.DependencyPath()),
		FileName:     "request",
	})
}

func prepareDirectoriesForRequestHelper(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{module, scaffold.TRANSPORT_PACKAGE, scaffold.REST_PACKAGE},
	)
}
