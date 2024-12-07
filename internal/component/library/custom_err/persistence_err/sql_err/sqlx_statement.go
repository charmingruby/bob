package sql_err

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/library/constant"
	"github.com/charmingruby/bob/internal/component/library/data"
	"github.com/charmingruby/bob/internal/component/shared/opt"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/scaffold"
)

func MakeSQLXStatementError(m filesystem.Manager) filesystem.File {
	return base.New(base.ComponentInput{
		Package: constant.SQL_ERROR_PACKAGE,
		DestinationDirectory: scaffold.CustomErrPath(
			m.ModuleDirectory(scaffold.SHARED_MODULE),
			[]string{constant.PERSISTENCE_ERR_PACKAGE, constant.SQL_ERROR_PACKAGE},
		),
	}).Componetize(base.ComponetizeInput{
		TemplateName: constant.SQL_STATEMENT_ERR_TEMPLATE,
		TemplateData: data.NewSQLXStatementErrData(m.DependencyPath(), opt.POSTGRES_DATABASE),
		FileName:     "sqlx_statement",
		FileSuffix:   "err",
	})
}
