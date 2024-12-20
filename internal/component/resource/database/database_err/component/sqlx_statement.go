package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/resource/database/database_err/constant"
	"github.com/charmingruby/bob/internal/component/resource/database/database_err/data"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeSQLXStatementError(m filesystem.Manager) filesystem.File {
	return base.New(base.ComponentInput{
		Package: constant.SQL_ERROR_PACKAGE,
		DestinationDirectory: definition.CustomErrPath(
			m.ModuleDirectory(definition.SHARED_MODULE),
			[]string{constant.PERSISTENCE_ERR_PACKAGE, constant.SQL_ERROR_PACKAGE},
		),
	}).Componetize(
		definition.ADD_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.SQL_STATEMENT_ERR_TEMPLATE,
			TemplateData: data.NewSQLXStatementErrData(m.DependencyPath(), "postgres"),
			FileName:     "sqlx_statement",
			FileSuffix:   "err",
		})
}
