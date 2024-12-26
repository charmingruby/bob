package postgres_db

import (
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/context/rest/bundle"
	"github.com/charmingruby/bob/internal/component/context/rest/module/postgres_db/component"
	"github.com/charmingruby/bob/internal/component/core/bundle/core"
	"github.com/charmingruby/bob/internal/component/shared/resource/database/postgres"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func PerformWithPostgresDatabase(m filesystem.Manager, module, modelName, tableName string) error {
	newModule := component.MakeRegistryWithPostgresDatabase(m, module, modelName)
	if err := m.GenerateFile(newModule); err != nil {
		return err
	}

	output.ComponentCreated(newModule.Identifier)

	if err := postgres.PerformPostgresRepository(m, module, modelName, tableName, true); err != nil {
		return err
	}

	if err := core.PerformCore(m, module, modelName); err != nil {
		return err
	}

	if err := bundle.PerformRest(m, module); err != nil {
		return err
	}

	return nil
}
