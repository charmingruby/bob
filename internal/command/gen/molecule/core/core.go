package core

import (
	"github.com/charmingruby/bob/internal/command/gen/atom"
	"github.com/charmingruby/bob/internal/command/gen/molecule/service"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
	"github.com/charmingruby/bob/internal/command/shared/scaffold"
)

func MakeCore(m filesystem.Manager, module string) {
	prepareDirectoriesForCore(m, module)

	sampleActor := module

	service.MakeService(m, sampleActor, module)

	if err := m.GenerateFile(atom.MakeRepositoryComponent(m, module, sampleActor)); err != nil {
		panic(err)
	}

	if err := m.GenerateFile(atom.MakeModelComponent(m, module, sampleActor)); err != nil {
		panic(err)
	}
}

func prepareDirectoriesForCore(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{scaffold.CORE_PACKAGE},
	)
}
