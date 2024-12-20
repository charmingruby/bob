package molecule

import (
	"github.com/charmingruby/bob/internal/component/atom"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/shared"
)

func PerformCore(m filesystem.Manager, module, baseModelName string) {
	prepareDirectoriesForCore(m, module)

	PerformService(m, baseModelName, module)

	components := []filesystem.File{
		atom.MakeModel(m, module, baseModelName),
		atom.MakeRepository(m, module, baseModelName),
	}

	for _, component := range components {
		if err := m.GenerateFile(component); err != nil {
			panic(err)
		}
	}
}

func prepareDirectoriesForCore(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{shared.CORE_PACKAGE},
	)
}
