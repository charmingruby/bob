package persistence_err

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/library/constant"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/scaffold"
)

func MakePersistenceError(m filesystem.Manager) filesystem.File {
	return base.New(base.ComponentInput{
		Package: constant.SQL_ERROR_PACKAGE,
		DestinationDirectory: scaffold.CustomErrPath(
			m.ModuleDirectory(scaffold.SHARED_MODULE),
			[]string{constant.PERSISTENCE_ERR_PACKAGE},
		),
	}).Componetize(base.ComponetizeInput{
		TemplateName: constant.PERSISTENCE_ERR_TEMPLATE,
		FileName:     "persistence",
		FileSuffix:   "err",
	})
}
