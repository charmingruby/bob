package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakePostgresConnection(m filesystem.Manager) filesystem.File {

	prepareDirectoriesForConnection(m)

	template := "resource/database/postgres/connection"

	return base.New(base.ComponentInput{
		Package:              definition.POSTGRES_PACKAGE,
		DestinationDirectory: m.ExternalLibraryDirectory(definition.POSTGRES_PACKAGE),
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			FileName:     "connection",
		})
}

func prepareDirectoriesForConnection(m filesystem.Manager) {
	m.GenerateNestedDirectories(
		m.MainDirectory(),
		[]string{definition.LIBRARY_PACKAGE, definition.POSTGRES_PACKAGE},
	)
}
