package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/resource/database/postgres/constant"
	"github.com/charmingruby/bob/internal/filesystem"
)

func MakePostgresConnection(m filesystem.Manager) filesystem.File {
	return base.New(base.ComponentInput{
		Package:              constant.POSTGRES_PACKAGE,
		DestinationDirectory: m.ExternalLibraryDirectory(constant.POSTGRES_PACKAGE),
	}).Componetize(base.ComponetizeInput{
		TemplateName: constant.POSTGRES_CONNECTION_TEMPLATE,
		FileName:     "connection",
	})
}
