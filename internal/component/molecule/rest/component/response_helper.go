package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/molecule/rest/constant"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/shared"
)

func MakeResponseHelper(m filesystem.Manager) filesystem.File {
	prepareDirectoriesForResponseHelper(m, shared.SHARED_MODULE)

	return base.New(base.ComponentInput{
		Package: shared.REST_PACKAGE,
		DestinationDirectory: shared.TransportPath(
			m.ModuleDirectory(shared.SHARED_MODULE),
			shared.REST_PACKAGE,
			nil,
		),
	}).Componetize(
		shared.GENERATE_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.REST_RESPONSE_HELPER_TEMPLATE,
			FileName:     "response",
		})
}

func prepareDirectoriesForResponseHelper(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{module, shared.TRANSPORT_PACKAGE, shared.REST_PACKAGE},
	)
}
