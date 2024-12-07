package resource

import (
	"github.com/charmingruby/bob/internal/component/resource/database/database_err"
	errConst "github.com/charmingruby/bob/internal/component/resource/database/database_err/constant"
	"github.com/charmingruby/bob/internal/component/resource/database/postgres/component"
	pgConst "github.com/charmingruby/bob/internal/component/resource/database/postgres/constant"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/scaffold"
)

func MakePostgresRepository(m filesystem.Manager, module, modelName string) filesystem.File {
	return component.MakePostgresRepository(m, module, modelName)
}

func MakePostgresDependencies(m filesystem.Manager) {
	prepareDirectoriesForPostgresDependencies(m)

	conn := component.MakePostgresConnection(m)
	if err := m.GenerateFile(conn); err != nil {
		panic(err)
	}

	persistenceErr := database_err.MakePersistenceError(m)
	if err := m.GenerateFile(persistenceErr); err != nil {
		panic(err)
	}

	sqlErr := database_err.MakeSQLXStatementError(m)
	if err := m.GenerateFile(sqlErr); err != nil {
		panic(err)
	}
}

func prepareDirectoriesForPostgresDependencies(m filesystem.Manager) {
	m.GenerateNestedDirectories(
		m.MainDirectory(),
		[]string{scaffold.LIBRARY_PACKAGE, pgConst.POSTGRES_PACKAGE},
	)

	m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{scaffold.SHARED_MODULE, errConst.CUSTOM_ERR_PACKAGE, errConst.PERSISTENCE_ERR_PACKAGE, errConst.SQL_ERROR_PACKAGE},
	)
}
