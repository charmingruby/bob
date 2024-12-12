package molecule

import (
	"github.com/charmingruby/bob/internal/component/atom"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/scaffold"
)

func MakeAndRunCore(m filesystem.Manager, module, database string) {
	prepareDirectoriesForCore(m, module)

	sampleActor := module

	MakeAndRunService(m, sampleActor, module)

	repository := atom.MakeRepository(m, module, sampleActor)
	if err := m.GenerateFile(repository); err != nil {
		panic(err)
	}

	model := atom.MakeModel(m, module, sampleActor)
	if err := m.GenerateFile(model); err != nil {
		panic(err)
	}
}

func prepareDirectoriesForCore(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{scaffold.CORE_PACKAGE},
	)
}
