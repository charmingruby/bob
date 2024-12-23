package bundle

import (
	"github.com/charmingruby/bob/internal/component/architecture/unit"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func PerformCore(m filesystem.Manager, module, baseModelName string) {
	prepareDirectoriesForCore(m, module)

	PerformService(m, baseModelName, module)

	components := []filesystem.File{
		unit.MakeModel(m, module, baseModelName),
		unit.MakeRepository(m, module, baseModelName),
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
		[]string{definition.CORE_PACKAGE},
	)
}
