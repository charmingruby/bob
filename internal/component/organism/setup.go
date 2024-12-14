package organism

import (
	"github.com/charmingruby/bob/internal/component/organism/setup/component"
	"github.com/charmingruby/bob/internal/filesystem"
)

func MakeAndRunSetup(m filesystem.Manager) {
	prepareDirectoriesForSetup(m)

	baseModule := "example"
	baseTableName := "examples"
	baseModelName := "example"

	entry := component.MakeEntry(m, baseModule, baseModelName)
	if err := m.GenerateFile(entry); err != nil {
		panic(err)
	}

	MakeAndRunModuleWithPostgresDatabase(m, baseModule, baseModelName, baseTableName)
}

func prepareDirectoriesForSetup(m filesystem.Manager) {
	m.GenerateNestedDirectories(
		m.MainDirectory(),
		[]string{"internal"},
	)
}
