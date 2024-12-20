package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/organism/bootstrap/constant"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeConfig(m filesystem.Manager) filesystem.File {
	prepareDirectoriesForConfig(m)

	return base.New(base.ComponentInput{
		DestinationDirectory: definition.RootPath([]string{"config"}),
	}).Componetize(
		definition.CREATE_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.CONFIG_TEMPLATE,
			FileName:     "config",
		})
}

func prepareDirectoriesForConfig(m filesystem.Manager) {
	m.GenerateNestedDirectories(
		m.MainDirectory(),
		[]string{"config"},
	)
}
