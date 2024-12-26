package core

import (
	"github.com/charmingruby/bob/internal/component/core/bundle/service"
	"github.com/charmingruby/bob/internal/component/core/unit"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func Perform(m filesystem.Manager, module, baseModelName string) ([]filesystem.File, error) {
	prepareDirectories(m, module)

	components := []filesystem.File{
		unit.MakeModel(m, module, baseModelName),
		unit.MakeRepository(m, module, baseModelName),
	}

	for _, component := range components {
		if err := m.GenerateFile(component); err != nil {
			return nil, err
		}
	}

	serviceComponents, err := service.Perfom(m, baseModelName, module)
	if err != nil {
		return nil, err
	}

	allComponents := append(components, serviceComponents...)

	return allComponents, nil
}

func prepareDirectories(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{definition.CORE_PACKAGE},
	)
}
