package resource

import (
	errComponent "github.com/charmingruby/bob/internal/component/resource/database/database_err/component"
	"github.com/charmingruby/bob/internal/component/resource/database/postgres/component"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func PerformPostgresRepository(m filesystem.Manager, module, modelName, tableName string, needDeps bool) {
	repo := component.MakePostgresRepository(m, module, modelName)
	if err := m.GenerateFile(repo); err != nil {
		panic(err)
	}

	if tableName != "" {
		PerformPostgresMigration(m, tableName)
	}

	if needDeps {
		PerformPostgresDependencies(m)
	}
}

func PerformPostgresDependencies(m filesystem.Manager) {
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

func PerformPostgresMigration(m filesystem.Manager, tableName string) {
	component.RunMigration(m, tableName)
}

func prepareDirectoriesForPostgresDependencies(m filesystem.Manager) {
	m.GenerateNestedDirectories(
		m.MainDirectory(),
		[]string{definition.LIBRARY_PACKAGE, definition.POSTGRES_PACKAGE},
	)

	m.GenerateNestedDirectories(
		m.MainDirectory(),
		[]string{definition.SQL_MIGRATION},
	)

	m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{definition.SHARED_MODULE, definition.CUSTOM_ERR_PACKAGE, definition.DATABASE_ERR_PACKAGE, definition.SQL_ERROR_PACKAGE},
	)
}
