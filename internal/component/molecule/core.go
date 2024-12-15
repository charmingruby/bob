package molecule

import (
	"github.com/charmingruby/bob/internal/component/atom"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/shared"
)

func MakeAndRunCore(m filesystem.Manager, module, baseModelName string) {
	prepareDirectoriesForCore(m, module)

	MakeAndRunService(m, baseModelName, module)

	repository := atom.MakeRepository(m, module, baseModelName)
	if err := m.GenerateFile(repository); err != nil {
		panic(err)
	}

	model := atom.MakeModel(m, module, baseModelName)
	if err := m.GenerateFile(model); err != nil {
		panic(err)
	}
}

func prepareDirectoriesForCore(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{shared.CORE_PACKAGE},
	)
}
