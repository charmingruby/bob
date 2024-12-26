package component

import (
	"github.com/charmingruby/bob/internal/component/context/rest"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeServer(m filesystem.Manager) filesystem.File {
	prepareDirectoriesForServer(m, definition.SHARED_MODULE)

	template := rest.TemplatePath("component/server")

	destination := definition.TransportPath(
		m.ModuleDirectory(definition.SHARED_MODULE),
		definition.REST_PACKAGE,
		nil,
	)

	content := "rest server"

	return base.New(base.ComponentInput{
		Identifier:           base.BuildIdentifier(definition.SHARED_MODULE, content, destination),
		Package:              definition.REST_PACKAGE,
		DestinationDirectory: destination,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			FileName:     "server",
		})
}

func MakeBaseServerMiddleware(m filesystem.Manager) filesystem.File {
	template := rest.TemplatePath("component/base_server_middleware")

	destination := definition.TransportPath(
		m.ModuleDirectory(definition.SHARED_MODULE),
		definition.REST_PACKAGE,
		nil,
	)

	content := "rest server middlewares"

	return base.New(base.ComponentInput{
		Identifier:           base.BuildIdentifier(definition.SHARED_MODULE, content, destination),
		Package:              definition.SERVICE_PACKAGE,
		DestinationDirectory: destination,
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
