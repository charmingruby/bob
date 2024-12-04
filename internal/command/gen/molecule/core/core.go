package core

import (
	"github.com/charmingruby/bob/internal/command/gen/atom"
	"github.com/charmingruby/bob/internal/command/gen/molecule/service"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

func CorePath() string {
	return "core"
}

func MakeCore(m filesystem.Manager, module string) {
	if err := m.GenerateDirectory(m.ModuleDirectory(module), "core"); err != nil {
		panic(err)
	}

	if err := m.GenerateMultipleDirectories(
		filesystem.ModulePath(m.SourceDirectory, module, CorePath()),
		[]string{"service", "model", "repository"},
	); err != nil {
		panic(err)
	}

	sampleActor := module

	service.MakeService(m, sampleActor, module)

	if err := m.GenerateFile(atom.MakeRepositoryComponent(m, module, sampleActor)); err != nil {
		panic(err)
	}

	if err := m.GenerateFile(atom.MakeModelComponent(m, module, sampleActor)); err != nil {
		panic(err)
	}
}
