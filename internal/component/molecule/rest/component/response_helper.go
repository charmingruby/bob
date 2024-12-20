package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/molecule/rest/constant"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeResponseHelper(m filesystem.Manager) filesystem.File {
	prepareDirectoriesForResponseHelper(m, definition.SHARED_MODULE)

	return base.New(base.ComponentInput{
		Package: definition.REST_PACKAGE,
		DestinationDirectory: definition.TransportPath(
			m.ModuleDirectory(definition.SHARED_MODULE),
			definition.REST_PACKAGE,
			nil,
		),
	}).Componetize(
		definition.ADD_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.REST_RESPONSE_HELPER_TEMPLATE,
			FileName:     "response",
		})
}

func prepareDirectoriesForResponseHelper(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{module, definition.TRANSPORT_PACKAGE, definition.REST_PACKAGE},
	)
}
