package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/molecule/rest/constant"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/shared"
)

func MakeServer(m filesystem.Manager) filesystem.File {
	prepareDirectoriesForServer(m, shared.SHARED_MODULE)

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
			TemplateName: constant.REST_SERVER,
			FileName:     "server",
		})
}

func MakeBaseServerMiddleware(m filesystem.Manager) filesystem.File {
	return base.New(base.ComponentInput{
		Package: shared.SERVICE_PACKAGE,
		DestinationDirectory: shared.TransportPath(
			m.ModuleDirectory(shared.SHARED_MODULE),
			shared.REST_PACKAGE,
			nil,
		),
	}).Componetize(
		shared.GENERATE_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.REST_BASE_SERVER_MIDDLEWARE,
			FileName:     "middleware",
		})
}

func prepareDirectoriesForServer(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{module, shared.TRANSPORT_PACKAGE, shared.REST_PACKAGE},
	)
}
