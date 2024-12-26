package postgres

import (
	"github.com/charmingruby/bob/internal/component/context/rest/module/postgres_db"
	"github.com/charmingruby/bob/internal/component/context/rest/template/postgres/component"
	sharedComponent "github.com/charmingruby/bob/internal/component/context/rest/template/shared/component"
	"github.com/charmingruby/bob/internal/component/shared/library"
	"github.com/charmingruby/bob/internal/component/shared/resource/container"
	"github.com/charmingruby/bob/internal/component/shared/resource/git"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

var (
	baseModule    = "example"
	baseModelName = "example"
)

func PerformWithPostgres(m filesystem.Manager, goVersion string) error {
	prepareDirectoriesForScaffold(m)

	baseTableName := "examples"

	components := []filesystem.File{
		sharedComponent.MakeGoMod(m, goVersion),
		component.MakeEntry(m, baseModule, baseModelName),
		component.MakeConfig(m),
		component.MakeEnvironmentExample(m),
		component.MakeMakefile(m),
		library.MakeAir(m),
	}

	for _, c := range components {
		if err := m.GenerateFile(c); err != nil {
			return err
		}
	}

	if err := container.PerformDockerContainer(m, goVersion); err != nil {
		return err
	}

	if err := container.PerformDockerComposeWithPostgres(m); err != nil {
		return err
	}

	if err := git.PerformGitignore(m); err != nil {
		return err
	}

	if err := postgres_db.PerformWithPostgresDatabase(m, baseModule, baseModelName, baseTableName); err != nil {
		return err
	}

	return nil
}

func prepareDirectoriesForScaffold(m filesystem.Manager) {
	m.GenerateNestedDirectories(
		m.MainDirectory(),
		[]string{"internal"},
	)
}
