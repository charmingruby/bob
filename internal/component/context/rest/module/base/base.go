package base

import (
	"github.com/charmingruby/bob/internal/component/context/rest/bundle/setup"
	"github.com/charmingruby/bob/internal/component/context/rest/module/base/component"
	"github.com/charmingruby/bob/internal/component/core/bundle/core"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func Perform(m filesystem.Manager, module, modelName string) ([]filesystem.File, error) {
	newModule := component.MakeRegistry(m, module)
	if err := m.GenerateFile(newModule); err != nil {
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

	allComponents := append([]filesystem.File{newModule}, coreComponents...)
	allComponents = append(allComponents, setupComponents...)

	return allComponents, nil
}
