package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeResponseHelper(m filesystem.Manager) filesystem.File {
	prepareDirectoriesForResponseHelper(m, definition.SHARED_MODULE)

	template := "architecture/bundle/rest/response_helper"

	return base.New(base.ComponentInput{
		Package: definition.REST_PACKAGE,
		DestinationDirectory: definition.TransportPath(
			m.ModuleDirectory(definition.SHARED_MODULE),
			definition.REST_PACKAGE,
			nil,
		),
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			FileName:     "response",
		})
}

func prepareDirectoriesForResponseHelper(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{module, definition.TRANSPORT_PACKAGE, definition.REST_PACKAGE},
	)
}
