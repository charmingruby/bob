package custom_db

import (
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/context/rest/bundle/setup"
	"github.com/charmingruby/bob/internal/component/context/rest/module/custom_db/component"
	"github.com/charmingruby/bob/internal/component/core/bundle/core"
	"github.com/charmingruby/bob/internal/component/core/unit"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func Perform(m filesystem.Manager, module, modelName, database string) error {
	newModule := component.MakeRegistry(m, module, modelName, database)
	if err := m.GenerateFile(newModule); err != nil {
		return err
	}

	output.ComponentCreated(newModule.Identifier)

	repo := unit.MakeUnimplementedRepository(m, module, modelName, database)
	if err := m.GenerateFile(repo); err != nil {
		return err
	}

	output.ComponentCreated(repo.Identifier)

	if err := core.Perform(m, module, modelName); err != nil {
		return err
	}

	if err := setup.Perform(m, module); err != nil {
		return err
	}

	return nil
}
