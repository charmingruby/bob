package core

import (
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/core/bundle/service"
	"github.com/charmingruby/bob/internal/component/core/unit"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func Perform(m filesystem.Manager, module, baseModelName string) error {
	prepareDirectories(m, module)

	components := []filesystem.File{
		unit.MakeModel(m, module, baseModelName),
		unit.MakeRepository(m, module, baseModelName),
	}

	for _, component := range components {
		if err := m.GenerateFile(component); err != nil {
			return err
		}

		output.ComponentCreated(component.Identifier)
	}

	if err := service.Perfom(m, baseModelName, module); err != nil {
		return err
	}

	return nil
}

func prepareDirectories(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{definition.CORE_PACKAGE},
	)
}
