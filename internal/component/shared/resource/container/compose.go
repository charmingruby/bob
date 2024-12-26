package container

import (
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeCompose(m filesystem.Manager) filesystem.File {
	template := TemplatePath("raw_compose")

	destination := m.MainDirectory()

	resource := "docker"

	content := "docker-compose with api"

	return base.New(base.ComponentInput{
		Identifier:           base.BuildNonModuleIdentifier(resource, content, destination),
		DestinationDirectory: m.RootDirectory,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			FileName:     "docker-compose",
			Extension:    definition.YML_EXTENSION,
		})
}
