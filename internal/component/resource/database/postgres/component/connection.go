package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/resource/database/postgres/constant"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakePostgresConnection(m filesystem.Manager) filesystem.File {
	prepareDirectoriesForConnection(m)

	return base.New(base.ComponentInput{
		Package:              constant.POSTGRES_PACKAGE,
		DestinationDirectory: m.ExternalLibraryDirectory(constant.POSTGRES_PACKAGE),
	}).Componetize(
		definition.ADD_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.POSTGRES_CONNECTION_TEMPLATE,
			FileName:     "connection",
		})
}

func prepareDirectoriesForConnection(m filesystem.Manager) {
	m.GenerateNestedDirectories(
		m.MainDirectory(),
		[]string{definition.LIBRARY_PACKAGE, constant.POSTGRES_PACKAGE},
	)
}
