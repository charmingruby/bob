package component

import (
	"github.com/charmingruby/bob/internal/component/resource/container/data"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeContainer(m filesystem.Manager, goVersion string) filesystem.File {
	template := "resource/container/container"

	destination := m.RootDirectory

	resource := "docker"

	content := "dockerfile"

	return base.New(base.ComponentInput{
		Identifier:           base.BuildNonModuleIdentifier(resource, content, destination),
		DestinationDirectory: m.RootDirectory,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			TemplateData: data.NewContainerData(goVersion),
			FileName:     "Dockerfile",
			Extension:    definition.NO_EXTENSION,
		})
}
