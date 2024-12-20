package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/organism/bootstrap/constant"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeCompose(m filesystem.Manager) filesystem.File {
	return base.New(base.ComponentInput{
		DestinationDirectory: m.RootDirectory,
	}).Componetize(
		definition.CREATE_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.COMPOSE_TEMPLATE,
			FileName:     "docker-compose",
			Extension:    definition.YML_EXTENSION,
		})
}
