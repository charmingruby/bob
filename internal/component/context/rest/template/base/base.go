package base

import (
	"github.com/charmingruby/bob/internal/component/context/rest/bundle/health_check"
	"github.com/charmingruby/bob/internal/component/context/rest/module/custom_db"
	"github.com/charmingruby/bob/internal/component/context/rest/template/base/component"
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

func Perfom(m filesystem.Manager, goVersion, database string) ([]filesystem.File, error) {
	prepareDirectoriesForScaffold(m)

	components := []filesystem.File{
		sharedComponent.MakeGoMod(m, goVersion),
		component.MakeEntry(m, baseModule, baseModelName),
		component.MakeConfig(m),
		component.MakeEnvironmentExample(m),
		component.MakeMakefile(m),
		docker.MakeContainer(m, goVersion),
		docker.MakeCompose(m),
		git.MakeGitIgnore(m),
		library.MakeAir(m),
	}

	for _, c := range components {
		if err := m.GenerateFile(c); err != nil {
			return nil, err
		}
	}

	dbComponents, err := custom_db.Perform(m, baseModule, baseModelName, database)
	if err != nil {
		return nil, err
	}

	healthCheckComponents, err := health_check.Perform(m)
	if err != nil {
		return nil, err
	}

	allComponents := append(components, dbComponents...)
	allComponents = append(allComponents, healthCheckComponents...)

	return allComponents, nil
}

func prepareDirectoriesForScaffold(m filesystem.Manager) {
	m.GenerateNestedDirectories(
		m.MainDirectory(),
		[]string{"internal"},
	)
}
