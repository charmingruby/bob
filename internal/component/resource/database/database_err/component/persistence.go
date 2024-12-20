package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/resource/database/database_err/constant"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakePersistenceError(m filesystem.Manager) filesystem.File {
	return base.New(base.ComponentInput{
		Package: constant.SQL_ERROR_PACKAGE,
		DestinationDirectory: definition.CustomErrPath(
			m.ModuleDirectory(definition.SHARED_MODULE),
			[]string{constant.PERSISTENCE_ERR_PACKAGE},
		),
	}).Componetize(
		definition.ADD_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.PERSISTENCE_ERR_TEMPLATE,
			FileName:     "persistence",
			FileSuffix:   "err",
		})
}
