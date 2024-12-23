package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeConfig(m filesystem.Manager) filesystem.File {
	prepareDirectoriesForConfig(m)

	template := "setup/scaffold/template/postgres/config"

	return base.New(base.ComponentInput{
		DestinationDirectory: definition.RootPath([]string{"config"}),
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
