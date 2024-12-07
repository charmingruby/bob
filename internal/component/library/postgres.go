package library

import (
	"github.com/charmingruby/bob/internal/component/library/constant"
	persistence_err "github.com/charmingruby/bob/internal/component/library/custom_err/database_err"
	"github.com/charmingruby/bob/internal/component/library/custom_err/database_err/sql_err"
	"github.com/charmingruby/bob/internal/component/library/postgres"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/scaffold"
)

func MakePostgresRepository(m filesystem.Manager, module, modelName string) filesystem.File {
	return postgres.MakePostgresRepository(m, module, modelName)
}

func MakePostgresDependencies(m filesystem.Manager) {
	prepareDirectoriesForPostgresDependencies(m)

	conn := postgres.MakePostgresConnection(m)
	if err := m.GenerateFile(conn); err != nil {
		panic(err)
	}

	persistenceErr := persistence_err.MakePersistenceError(m)
	if err := m.GenerateFile(persistenceErr); err != nil {
		panic(err)
	}

	sqlErr := sql_err.MakeSQLXStatementError(m)
	if err := m.GenerateFile(sqlErr); err != nil {
		panic(err)
	}
}

func prepareDirectoriesForPostgresDependencies(m filesystem.Manager) {
	m.GenerateNestedDirectories(
		m.MainDirectory(),
		[]string{scaffold.LIBRARY_PACKAGE, constant.POSTGRES_PACKAGE},
	)

	m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{scaffold.SHARED_MODULE, constant.CUSTOM_ERR_PACKAGE, constant.PERSISTENCE_ERR_PACKAGE, constant.SQL_ERROR_PACKAGE},
	)
}
