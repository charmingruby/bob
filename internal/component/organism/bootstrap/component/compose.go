package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/organism/bootstrap/constant"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/shared"
)

func MakeCompose(m filesystem.Manager) filesystem.File {
	return base.New(base.ComponentInput{
		DestinationDirectory: m.RootDirectory,
	}).Componetize(
		shared.CREATE_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.COMPOSE_TEMPLATE,
			FileName:     "docker-compose",
			Extension:    shared.YML_EXTENSION,
		})
}
