package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/resource/database/database_err/constant"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/shared"
)

func MakePersistenceError(m filesystem.Manager) filesystem.File {
	return base.New(base.ComponentInput{
		Package: constant.SQL_ERROR_PACKAGE,
		DestinationDirectory: shared.CustomErrPath(
			m.ModuleDirectory(shared.SHARED_MODULE),
			[]string{constant.PERSISTENCE_ERR_PACKAGE},
		),
	}).Componetize(
		shared.GENERATE_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.PERSISTENCE_ERR_TEMPLATE,
			FileName:     "persistence",
			FileSuffix:   "err",
		})
}
