package custom_db

import (
	"github.com/charmingruby/bob/internal/component/context/rest/bundle/setup"
	"github.com/charmingruby/bob/internal/component/context/rest/module/custom_db/component"
	"github.com/charmingruby/bob/internal/component/core/bundle/core"
	"github.com/charmingruby/bob/internal/component/core/unit"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func Perform(m filesystem.Manager, module, modelName, database string) ([]filesystem.File, error) {
	newModule := component.MakeRegistry(m, module, modelName, database)
	if err := m.GenerateFile(newModule); err != nil {
		return nil, err
	}

	repo := unit.MakeUnimplementedRepository(m, module, modelName, database)
	if err := m.GenerateFile(repo); err != nil {
		return nil, err
	}

	coreComponents, err := core.Perform(m, module, modelName)
	if err != nil {
		return nil, err
	}

	setupComponents, err := setup.Perform(m, module)
	if err != nil {
		return nil, err
	}

	allComponents := append([]filesystem.File{newModule, repo}, coreComponents...)
	allComponents = append(allComponents, setupComponents...)

	return allComponents, nil
}
