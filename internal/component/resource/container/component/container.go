package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/resource/container/data"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeContainer(m filesystem.Manager, goVersion string) filesystem.File {
	template := "resource/container/container"

	return base.New(base.ComponentInput{
		DestinationDirectory: m.RootDirectory,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			TemplateData: data.NewContainerData(goVersion),
			FileName:     "Dockerfile",
			Extension:    definition.NO_EXTENSION,
		})
}
