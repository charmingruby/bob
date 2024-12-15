package resource

import (
	errComponent "github.com/charmingruby/bob/internal/component/resource/database/database_err/component"
	errConst "github.com/charmingruby/bob/internal/component/resource/database/database_err/constant"
	"github.com/charmingruby/bob/internal/component/resource/database/postgres/component"
	pgConst "github.com/charmingruby/bob/internal/component/resource/database/postgres/constant"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/shared"
)

func MakeAndRunPostgresRepository(m filesystem.Manager, module, modelName, tableName string, needDeps bool) filesystem.File {
	repo := component.MakePostgresRepository(m, module, modelName)
	if err := m.GenerateFile(repo); err != nil {
		panic(err)
	}

	if tableName != "" {
		MakeAndRunPostgresMigration(m, tableName)
	}

	if needDeps {
		MakeAndRunPostgresDependencies(m)
	}

	return component.MakePostgresRepository(m, module, modelName)
}

func MakeAndRunPostgresDependencies(m filesystem.Manager) {
	prepareDirectoriesForPostgresDependencies(m)

	conn := component.MakePostgresConnection(m)
	if err := m.GenerateFile(conn); err != nil {
		panic(err)
	}

	persistenceErr := errComponent.MakePersistenceError(m)
	if err := m.GenerateFile(persistenceErr); err != nil {
		panic(err)
	}

	sqlErr := errComponent.MakeSQLXStatementError(m)
	if err := m.GenerateFile(sqlErr); err != nil {
		panic(err)
	}
}

func MakeAndRunPostgresMigration(m filesystem.Manager, tableName string) {
	component.RunMigration(m, tableName)
}

func prepareDirectoriesForPostgresDependencies(m filesystem.Manager) {
	m.GenerateNestedDirectories(
		m.MainDirectory(),
		[]string{shared.LIBRARY_PACKAGE, pgConst.POSTGRES_PACKAGE},
	)

	m.GenerateNestedDirectories(
		m.MainDirectory(),
		[]string{pgConst.MIGRATIONS_DIR},
	)

	m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{shared.SHARED_MODULE, errConst.CUSTOM_ERR_PACKAGE, errConst.PERSISTENCE_ERR_PACKAGE, errConst.SQL_ERROR_PACKAGE},
	)
}
