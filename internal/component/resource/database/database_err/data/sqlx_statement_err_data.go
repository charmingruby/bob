package data

import "github.com/charmingruby/bob/internal/shared/definition/component/base"

type SQLXStatementErrData struct {
	SourcePath  string
	SQLDatabase string
}

func NewSQLXStatementErrData(sourcePath, database string) SQLXStatementErrData {
	return SQLXStatementErrData{
		SourcePath:  sourcePath,
		SQLDatabase: base.PrivateNameFormat(database),
	}
}
