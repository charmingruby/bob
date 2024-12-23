package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/setup/configure/constant"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeConfigure(m filesystem.Manager) filesystem.File {
	return base.New(base.ComponentInput{
		DestinationDirectory: m.RootDirectory,
	}).Componetize(
		definition.CREATE_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.CONFIGURE_TEMPLATE,
			FileName:     "bob",
			Extension:    definition.YML_EXTENSION,
		})
}
