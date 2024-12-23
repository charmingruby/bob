package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeComposeWithPostgres(m filesystem.Manager) filesystem.File {
	template := "resource/container/compose_with_pg"

	return base.New(base.ComponentInput{
		DestinationDirectory: m.RootDirectory,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			FileName:     "docker-compose",
			Extension:    definition.YML_EXTENSION,
		})
}
