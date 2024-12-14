package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/molecule/constant"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/scaffold"
)

func MakeResponseHelper(m filesystem.Manager) filesystem.File {
	prepareDirectoriesForResponseHelper(m, scaffold.SHARED_MODULE)

	return base.New(base.ComponentInput{
		Package: scaffold.REST_PACKAGE,
		DestinationDirectory: scaffold.TransportPath(
			m.ModuleDirectory(scaffold.SHARED_MODULE),
			scaffold.REST_PACKAGE,
			nil,
		),
	}).Componetize(
		scaffold.GENERATE_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.REST_RESPONSE_HELPER_TEMPLATE,
			FileName:     "response",
		})
}

func prepareDirectoriesForResponseHelper(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{module, scaffold.TRANSPORT_PACKAGE, scaffold.REST_PACKAGE},
	)
}
