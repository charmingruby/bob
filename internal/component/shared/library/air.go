package library

import (
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeAir(m filesystem.Manager) filesystem.File {
	template := "library/air"

	destination := m.RootDirectory

	resource := "air"

	content := "hot reload"

	return base.New(base.ComponentInput{
		Identifier:           base.BuildNonModuleIdentifier(resource, content, destination),
		DestinationDirectory: destination,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			FileName:     ".air",
			Extension:    definition.TOML_EXTENSION,
		})
}
