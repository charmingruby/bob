package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/resource/database/database_err/data"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeSQLXStatementError(m filesystem.Manager) filesystem.File {
	template := "error/database/sql/statement"

	return base.New(base.ComponentInput{
		Package: definition.SQL_ERROR_PACKAGE,
		DestinationDirectory: definition.CustomErrPath(
			m.ModuleDirectory(definition.SHARED_MODULE),
			[]string{definition.DATABASE_ERR_PACKAGE, definition.SQL_ERROR_PACKAGE},
		),
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			TemplateData: data.NewSQLXStatementErrData(m.DependencyPath(), "postgres"),
			FileName:     "sqlx_statement",
			FileSuffix:   "err",
		})
}
