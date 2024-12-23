package setup

import (
	"github.com/charmingruby/bob/internal/component/architecture/structure"
	"github.com/charmingruby/bob/internal/component/resource"
	"github.com/charmingruby/bob/internal/component/setup/scaffold/component"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func PerformScaffold(m filesystem.Manager, goVersion string) {
	prepareDirectoriesForScaffold(m)

	baseModule := "example"
	baseTableName := "examples"
	baseModelName := "example"

	components := []filesystem.File{
		component.MakeEntry(m, baseModule, baseModelName),
		component.MakeGoMod(m, goVersion),
		component.MakeConfig(m),
		component.MakeEnvironmentExample(m),
		component.MakeMakefile(m),
	}

	for _, c := range components {
		if err := m.GenerateFile(c); err != nil {
			panic(err)
		}
	}

	resource.PerformDockerContainer(m, goVersion)
	resource.PerformDockerCompose(m)
	resource.PerformGitignore(m)

	structure.PerformModuleWithPostgresDatabase(m, baseModule, baseModelName, baseTableName)
}

func prepareDirectoriesForScaffold(m filesystem.Manager) {
	m.GenerateNestedDirectories(
		m.MainDirectory(),
		[]string{"internal"},
	)
}
