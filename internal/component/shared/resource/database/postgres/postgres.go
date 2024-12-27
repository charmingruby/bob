package postgres

import (
	errComponent "github.com/charmingruby/bob/internal/component/shared/err/database/component"
	"github.com/charmingruby/bob/internal/component/shared/resource/database/postgres/component"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func PerformRepository(m filesystem.Manager, module, modelName, tableName string, needDeps bool) ([]filesystem.File, error) {
	repo := component.MakePostgresRepository(m, module, modelName)
	if err := m.GenerateFile(repo); err != nil {
		return nil, err
	}

	components := []filesystem.File{repo}

	if tableName != "" {
		migrationComponents, err := PerformMigration(m, tableName)
		if err != nil {
			return nil, err
		}

		components = append(components, migrationComponents...)
	}

	if needDeps {
		depsComponents, err := PerformDependencies(m)
		if err != nil {
			return nil, err
		}

		components = append(components, depsComponents...)
	}

	return components, nil
}

func PerformDependencies(m filesystem.Manager) ([]filesystem.File, error) {
	prepareDirectoriesForDependencies(m)

	components := []filesystem.File{
		component.MakePostgresConnection(m),
		errComponent.MakePersistenceError(m),
		errComponent.MakeSQLXStatementError(m),
	}

	for _, f := range components {
		if err := m.GenerateFile(f); err != nil {
			return nil, err
		}
	}

	return components, nil
}

func PerformMigration(m filesystem.Manager, tableName string) ([]filesystem.File, error) {
	return component.RunMigration(m, tableName)
}

func prepareDirectoriesForDependencies(m filesystem.Manager) {
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
