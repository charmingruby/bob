package service

import (
	"github.com/charmingruby/bob/internal/component/core/bundle/service/component"
	"github.com/charmingruby/bob/internal/component/core/unit"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/charmingruby/bob/pkg/util"
)

func Perfom(m filesystem.Manager, repo string, module string) ([]filesystem.File, error) {
	prepareDirectories(m, module)

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
			return nil, err
		}
	}

	return components, nil
}

func prepareDirectories(m filesystem.Manager, module string) {
	if err := m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{"core", "service"},
	); err != nil {
		panic(err)
	}
}
