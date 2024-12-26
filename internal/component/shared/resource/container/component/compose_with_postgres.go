package component

import (
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeComposeWithPostgres(m filesystem.Manager) filesystem.File {
	template := "resource/container/compose_with_pg"

	destination := m.MainDirectory()

	resource := "docker"

	content := "docker-compose with postgres"

	return base.New(base.ComponentInput{
		Identifier:           base.BuildNonModuleIdentifier(resource, content, destination),
		DestinationDirectory: destination,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			FileName:     "docker-compose",
			Extension:    definition.YML_EXTENSION,
		})
}
