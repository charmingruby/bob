package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/molecule"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/scaffold"
)

func MakeServerComponent(m filesystem.Manager) filesystem.File {
	prepareDirectoriesForServer(m, scaffold.SHARED_MODULE)

	return base.New(base.ComponentInput{
		Package: scaffold.REST_PACKAGE,
		DestinationDirectory: scaffold.TransportPath(
			m.ModuleDirectory(scaffold.SHARED_MODULE),
			scaffold.REST_PACKAGE,
			nil,
		),
	}).Componetize(base.ComponetizeInput{
		TemplateName: molecule.REST_SERVER,
		FileName:     "server",
	})
}

func MakeBaseServerMiddlewareComponent(m filesystem.Manager) filesystem.File {
	return base.New(base.ComponentInput{
		Package: scaffold.SERVICE_PACKAGE,
		DestinationDirectory: scaffold.TransportPath(
			m.ModuleDirectory(scaffold.SHARED_MODULE),
			scaffold.REST_PACKAGE,
			nil,
		),
	}).Componetize(base.ComponetizeInput{
		TemplateName: molecule.REST_BASE_SERVER_MIDDLEWARE,
		FileName:     "middleware",
	})
}

func prepareDirectoriesForServer(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{module, scaffold.TRANSPORT_PACKAGE, scaffold.REST_PACKAGE},
	)
}
