package postgres

import (
	"github.com/charmingruby/bob/internal/component/context/rest/bundle/health_check"
	"github.com/charmingruby/bob/internal/component/context/rest/module/postgres"
	"github.com/charmingruby/bob/internal/component/context/rest/template/postgres/component"
	sharedComponent "github.com/charmingruby/bob/internal/component/context/rest/template/shared/component"
	"github.com/charmingruby/bob/internal/component/shared/library"
	"github.com/charmingruby/bob/internal/component/shared/resource/docker"
	"github.com/charmingruby/bob/internal/component/shared/resource/git"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

var (
	baseModule    = "example"
	baseModelName = "example"
)

func PerformWithPostgres(m filesystem.Manager, goVersion string) ([]filesystem.File, error) {
	prepareDirectoriesForScaffold(m)

	baseTableName := "examples"

	components := []filesystem.File{
		sharedComponent.MakeGoMod(m, goVersion),
		component.MakeEntry(m, baseModule, baseModelName),
		component.MakeConfig(m),
		component.MakeEnvironmentExample(m),
		component.MakeMakefile(m),
		docker.MakeContainer(m, goVersion),
		docker.MakeComposeWithPostgres(m),
		git.MakeGitIgnore(m),
		library.MakeAir(m),
	}

	for _, c := range components {
		if err := m.GenerateFile(c); err != nil {
			return nil, err
		}
	}

	postgresComponents, err := postgres.Perform(m, baseModule, baseModelName, baseTableName)
	if err != nil {
		return nil, err
	}

	healthCheckComponents, err := health_check.Perform(m)
	if err != nil {
		return nil, err
	}

	allComponents := append(components, postgresComponents...)
	allComponents = append(allComponents, healthCheckComponents...)

	return allComponents, nil
}

func prepareDirectoriesForScaffold(m filesystem.Manager) {
	m.GenerateNestedDirectories(
		m.MainDirectory(),
		[]string{"internal"},
	)
}
