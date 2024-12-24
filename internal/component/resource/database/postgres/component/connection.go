package component

import (
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakePostgresConnection(m filesystem.Manager) filesystem.File {
	prepareDirectoriesForConnection(m)

	template := "resource/database/postgres/connection"

	destination := m.ExternalLibraryDirectory(definition.POSTGRES_PACKAGE)

	content := "postgres connection"

	return base.New(base.ComponentInput{
		Identifier:           base.BuildIdentifier(definition.LIBRARY_PACKAGE, content, destination),
		Package:              definition.POSTGRES_PACKAGE,
		DestinationDirectory: destination,
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
