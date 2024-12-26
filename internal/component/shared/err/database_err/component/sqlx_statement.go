package component

import (
	"github.com/charmingruby/bob/internal/component/shared/err"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

type sQLXStatementErrData struct {
	SourcePath  string
	SQLDatabase string
}

func newSQLXStatementErrData(sourcePath, database string) sQLXStatementErrData {
	return sQLXStatementErrData{
		SourcePath:  sourcePath,
		SQLDatabase: base.PrivateNameFormat(database),
	}
}

func MakeSQLXStatementError(m filesystem.Manager) filesystem.File {
	template := err.TemplatePath("database_err/sql/statement")

	destination := definition.CustomErrPath(
		m.ModuleDirectory(definition.SHARED_MODULE),
		[]string{definition.DATABASE_ERR_PACKAGE, definition.SQL_ERROR_PACKAGE},
	)

	content := "sql statement error"

	return base.New(base.ComponentInput{
		Identifier:           base.BuildIdentifier(definition.SHARED_MODULE, content, destination),
		Package:              definition.SQL_ERROR_PACKAGE,
		DestinationDirectory: destination,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			TemplateData: newSQLXStatementErrData(m.DependencyPath(), "postgres"),
			FileName:     "sqlx_statement",
			FileSuffix:   "err",
		})
}
