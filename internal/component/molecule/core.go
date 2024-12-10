package molecule

import (
	"github.com/charmingruby/bob/internal/component/atom"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/scaffold"
)

func MakeCore(m filesystem.Manager, module, database string) {
	prepareDirectoriesForCore(m, module)

	sampleActor := module

	MakeService(m, sampleActor, module)

	repository := atom.MakeRepository(m, module, sampleActor, database)
	if err := m.GenerateFile(repository); err != nil {
		panic(err)
	}

	if database != "" {
		persistenceRepository := atom.MakePersistenceRepository(m, module, sampleActor, database)
		if err := m.GenerateFile(persistenceRepository); err != nil {
			panic(err)
		}
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
