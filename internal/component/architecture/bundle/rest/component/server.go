package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeServer(m filesystem.Manager) filesystem.File {
	prepareDirectoriesForServer(m, definition.SHARED_MODULE)

	template := "architecture/bundle/rest/server"

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
			FileName:     "server",
		})
}

func MakeBaseServerMiddleware(m filesystem.Manager) filesystem.File {
	template := "architecture/bundle/rest/base_server_middleware"

	return base.New(base.ComponentInput{
		Package: definition.SERVICE_PACKAGE,
		DestinationDirectory: definition.TransportPath(
			m.ModuleDirectory(definition.SHARED_MODULE),
			definition.REST_PACKAGE,
			nil,
		),
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			FileName:     "middleware",
		})
}

func prepareDirectoriesForServer(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{module, definition.TRANSPORT_PACKAGE, definition.REST_PACKAGE},
	)
}
