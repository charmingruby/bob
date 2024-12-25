package resource

import (
	"github.com/charmingruby/bob/internal/cli/output"
	errComponent "github.com/charmingruby/bob/internal/component/resource/database/database_err/component"
	"github.com/charmingruby/bob/internal/component/resource/database/postgres/component"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func PerformPostgresRepository(m filesystem.Manager, module, modelName, tableName string, needDeps bool) error {
	repo := component.MakePostgresRepository(m, module, modelName)
	if err := m.GenerateFile(repo); err != nil {
		return err
	}

	output.ComponentCreated(repo.Identifier)

	if tableName != "" {
		if err := PerformPostgresMigration(m, tableName); err != nil {
			return err
		}
	}

	if needDeps {
		return PerformPostgresDependencies(m)
	}

	return nil
}

func PerformPostgresDependencies(m filesystem.Manager) error {
	prepareDirectoriesForPostgresDependencies(m)

	components := []filesystem.File{
		component.MakePostgresConnection(m),
		errComponent.MakePersistenceError(m),
		errComponent.MakeSQLXStatementError(m),
	}

	for _, f := range components {
		if err := m.GenerateFile(f); err != nil {
			return err
		}

		output.ComponentCreated(f.Identifier)
	}

	return nil
}

func PerformPostgresMigration(m filesystem.Manager, tableName string) error {
	return component.RunMigration(m, tableName)
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
