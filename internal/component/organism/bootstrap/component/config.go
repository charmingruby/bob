package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/organism/bootstrap/constant"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/shared"
)

func MakeConfig(m filesystem.Manager) filesystem.File {
	prepareDirectoriesForConfig(m)

	return base.New(base.ComponentInput{
		DestinationDirectory: shared.RootPath([]string{"config"}),
	}).Componetize(
		shared.BOOTSTRAP_COMMAND,
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
