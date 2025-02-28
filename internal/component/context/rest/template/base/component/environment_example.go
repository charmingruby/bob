package component

import (
	"github.com/charmingruby/bob/internal/component/context/rest"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeEnvironmentExample(m filesystem.Manager) filesystem.File {
	template := rest.TemplatePath("template/base/environment_example")

	directory := m.MainDirectory()

	resource := "env"

	content := "environment variables example"

	return base.New(base.ComponentInput{
		Identifier:           base.BuildNonModuleIdentifier(resource, content, directory),
		DestinationDirectory: directory,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			FileName:     ".env.example",
			Extension:    definition.NO_EXTENSION,
		})
}
