package service

import (
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/core/bundle/service/component"
	"github.com/charmingruby/bob/internal/component/core/unit"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/charmingruby/bob/pkg/util"
)

func PerformService(m filesystem.Manager, repo string, module string) error {
	prepareDirectoriesForService(m, module)

	sampleActor := module

	registry := util.Ternary[filesystem.File](
		repo == "",
		component.MakeIndependentServiceRegistry(m, module),
		component.MakeServiceRegistry(m, module, repo),
	)

	components := []filesystem.File{
		registry,
		unit.MakeService(m, module, sampleActor),
	}

	for _, component := range components {
		if err := m.GenerateFile(component); err != nil {
			return err
		}

		output.ComponentCreated(component.Identifier)
	}

	return nil
}

func prepareDirectoriesForService(m filesystem.Manager, module string) {
	if err := m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{"core", "service"},
	); err != nil {
		panic(err)
	}
}
