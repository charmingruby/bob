package service

import (
	"github.com/charmingruby/bob/internal/component/atom"
	"github.com/charmingruby/bob/internal/component/molecule/service/component"
	"github.com/charmingruby/bob/internal/filesystem"
)

func MakeService(m filesystem.Manager, repo string, module string) {
	sampleActor := module

	service := atom.MakeServiceComponent(m, module, sampleActor)

	if err := m.GenerateFile(service); err != nil {
		panic(err)
	}

	if repo == "" {
		if err := m.GenerateFile(component.MakeIndependentServiceRegistryComponent(
			m,
			module,
		)); err != nil {
			panic(err)
		}
	} else {
		if err := m.GenerateFile(component.MakeServiceRegistryComponent(
			m,
			module,
			repo,
		)); err != nil {
			panic(err)
		}
	}
}
