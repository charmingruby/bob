package service

import (
	"github.com/charmingruby/bob/internal/command/gen/atom"
	"github.com/charmingruby/bob/internal/command/gen/molecule/service/service_component"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

func ServicePath() string {
	return "core/service"
}

func MakeService(m filesystem.Manager, repo string, module string) {
	hasRepo := repo != ""

	if !hasRepo {
		if err := m.GenerateFile(service_component.MakeIndependentServiceRegistryComponent(
			m,
			module,
		)); err != nil {
			panic(err)
		}
	} else {
		if err := m.GenerateFile(service_component.MakeServiceRegistryComponent(
			m,
			module,
			repo,
		)); err != nil {
			panic(err)
		}
	}

	sampleActor := module

	if err := m.GenerateFile(atom.MakeServiceComponent(m.SourceDirectory, module, sampleActor)); err != nil {
		panic(err)
	}
}
