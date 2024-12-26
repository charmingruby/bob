package custom_db

import (
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/context/rest/bundle"
	"github.com/charmingruby/bob/internal/component/context/rest/module/custom_db/component"
	"github.com/charmingruby/bob/internal/component/core/bundle/core"
	"github.com/charmingruby/bob/internal/component/core/unit"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func PerformWithCustomDatabase(m filesystem.Manager, module, modelName, database string) error {
	newModule := component.MakeRegistryWithCustomDatabase(m, module, modelName, database)
	if err := m.GenerateFile(newModule); err != nil {
		panic(err)
	}

	output.ComponentCreated(newModule.Identifier)

	repo := unit.MakeUnimplementedRepository(m, module, modelName, database)
	if err := m.GenerateFile(repo); err != nil {
		return err
	}

	output.ComponentCreated(repo.Identifier)

	if err := core.PerformCore(m, module, modelName); err != nil {
		return err
	}

	if err := bundle.PerformRest(m, module); err != nil {
		return err
	}

	return nil
}