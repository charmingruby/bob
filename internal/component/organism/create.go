package organism

import (
	"github.com/charmingruby/bob/internal/component/organism/bootstrap/component"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func PerformSetup(m filesystem.Manager, goVersion string) {
	prepareDirectoriesForSetup(m)

	baseModule := "example"
	baseTableName := "examples"
	baseModelName := "example"

	components := []filesystem.File{
		component.MakeEntry(m, baseModule, baseModelName),
		component.MakeGoMod(m, goVersion),
		component.MakeConfig(m),
		component.MakeEnvironmentExample(m),
		component.MakeCompose(m),
		component.MakeContainer(m, goVersion),
		component.MakeMakefile(m),
		component.MakeGitIgnore(m),
	}

	for _, c := range components {
		if err := m.GenerateFile(c); err != nil {
			panic(err)
		}
	}

	PerformModuleWithPostgresDatabase(m, baseModule, baseModelName, baseTableName)
}

func prepareDirectoriesForSetup(m filesystem.Manager) {
	m.GenerateNestedDirectories(
		m.MainDirectory(),
		[]string{"internal"},
	)
}
