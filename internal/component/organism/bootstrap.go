package organism

import (
	"github.com/charmingruby/bob/internal/component/organism/bootstrap/component"
	"github.com/charmingruby/bob/internal/filesystem"
)

func MakeAndRunSetup(m filesystem.Manager) {
	prepareDirectoriesForSetup(m)

	baseModule := "example"
	baseTableName := "examples"
	baseModelName := "example"

	components := []filesystem.File{
		component.MakeEntry(m, baseModule, baseModelName),
		component.MakeGoMod(m, "1.23.3"),
		component.MakeConfig(m),
	}

	for _, c := range components {
		if err := m.GenerateFile(c); err != nil {
			panic(err)
		}
	}

	MakeAndRunModuleWithPostgresDatabase(m, baseModule, baseModelName, baseTableName)
}

func prepareDirectoriesForSetup(m filesystem.Manager) {
	m.GenerateNestedDirectories(
		m.MainDirectory(),
		[]string{"internal"},
	)
}
