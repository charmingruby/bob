package service

import (
	"github.com/charmingruby/bob/internal/command/gen/atom"
	"github.com/charmingruby/bob/internal/command/gen/molecule/service/service_component"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

func MakeService(m filesystem.Manager, repo string, module string) {
	sampleActor := module

	service := atom.MakeServiceComponent(m, module, sampleActor)

	if err := m.GenerateFile(service); err != nil {
		panic(err)
	}

	if repo == "" {
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
}
