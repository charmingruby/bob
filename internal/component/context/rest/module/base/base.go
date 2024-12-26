package base

import (
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/context/rest/bundle"
	"github.com/charmingruby/bob/internal/component/context/rest/module/base/component"
	"github.com/charmingruby/bob/internal/component/core/bundle/core"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func PerformBase(m filesystem.Manager, module, modelName string) error {
	newModule := component.MakeBaseRegistry(m, module)
	if err := m.GenerateFile(newModule); err != nil {
		return err
	}

	output.ComponentCreated(newModule.Identifier)

	if err := core.PerformCore(m, module, modelName); err != nil {
		return err
	}

	if err := bundle.PerformRest(m, module); err != nil {
		return err
	}

	return nil
}
