package core

import (
	"github.com/charmingruby/bob/internal/component/core/bundle/service"
	"github.com/charmingruby/bob/internal/component/core/unit"
	"github.com/charmingruby/bob/internal/component/shared/err/core_err/component"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func Perform(m filesystem.Manager, module, baseModelName string) ([]filesystem.File, error) {
	prepareDirectories(m, module)

	baseServiceName := "greeting"

	components := []filesystem.File{
		unit.MakeModel(m, module, baseModelName),
		unit.MakeRepository(m, module, baseModelName),
		component.MakeServiceError(m),
	}

	for _, component := range components {
		if err := m.GenerateFile(component); err != nil {
			return nil, err
		}
	}

	serviceComponents, err := service.Perfom(m, baseModelName, module, baseServiceName, baseModelName)
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

	m.GenerateNestedDirectories(
		m.ModuleDirectory(definition.SHARED_MODULE),
		[]string{definition.CUSTOM_ERR_PACKAGE, definition.CORE_ERR_PACKAGE},
	)
}
