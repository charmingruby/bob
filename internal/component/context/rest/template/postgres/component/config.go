package component

import (
	"github.com/charmingruby/bob/internal/component/context/rest"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeConfig(m filesystem.Manager) filesystem.File {
	prepareDirectoriesForConfig(m)

	template := rest.TemplatePath("template/postgres/config")

	destination := definition.RootPath([]string{"config"})

	resource := "env"

	content := "environment variables loader"

	return base.New(base.ComponentInput{
		Identifier:           base.BuildNonModuleIdentifier(resource, content, destination),
		DestinationDirectory: destination,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			FileName:     "config",
		})
}

func prepareDirectoriesForConfig(m filesystem.Manager) {
	m.GenerateNestedDirectories(
		m.MainDirectory(),
		[]string{"config"},
	)
}
