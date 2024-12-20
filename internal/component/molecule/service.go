package molecule

import (
	"github.com/charmingruby/bob/internal/component/atom"
	"github.com/charmingruby/bob/internal/component/molecule/service/component"
	"github.com/charmingruby/bob/internal/filesystem"
)

func PerformService(m filesystem.Manager, repo string, module string) {
	prepareDirectoriesForService(m, module)

	sampleActor := module

	service := atom.MakeService(m, module, sampleActor)

	if err := m.GenerateFile(service); err != nil {
		panic(err)
	}

	if repo == "" {
		if err := m.GenerateFile(component.MakeIndependentServiceRegistry(
			m,
			module,
		)); err != nil {
			panic(err)
		}
	} else {
		if err := m.GenerateFile(component.MakeServiceRegistry(
			m,
			module,
			repo,
		)); err != nil {
			panic(err)
		}
	}
}

func prepareDirectoriesForService(m filesystem.Manager, module string) {
	if err := m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{"core", "service"},
	); err != nil {
		panic(err)
	}
}
