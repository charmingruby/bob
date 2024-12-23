package setup

import (
	"github.com/charmingruby/bob/internal/component/architecture/structure"
	"github.com/charmingruby/bob/internal/component/resource"
	pgTemplateComponent "github.com/charmingruby/bob/internal/component/setup/scaffold/template/postgres/component"

	baseTemplateComponent "github.com/charmingruby/bob/internal/component/setup/scaffold/template/base/component"
	sharedComponent "github.com/charmingruby/bob/internal/component/setup/scaffold/template/shared/component"

	"github.com/charmingruby/bob/internal/shared/filesystem"
)

var (
	baseModule    = "example"
	baseModelName = "example"
)

func PerformPostgresTemplate(m filesystem.Manager, goVersion string) error {
	prepareDirectoriesForScaffold(m)

	baseTableName := "examples"

	components := []filesystem.File{
		sharedComponent.MakeGoMod(m, goVersion),
		pgTemplateComponent.MakeEntry(m, baseModule, baseModelName),
		pgTemplateComponent.MakeConfig(m),
		pgTemplateComponent.MakeEnvironmentExample(m),
		pgTemplateComponent.MakeMakefile(m),
	}

	for _, c := range components {
		if err := m.GenerateFile(c); err != nil {
			return err
		}
	}

	if err := resource.PerformDockerContainer(m, goVersion); err != nil {
		return err
	}

	if err := resource.PerformDockerComposeWithPostgres(m); err != nil {
		return err
	}

	if err := resource.PerformGitignore(m); err != nil {
		return err
	}

	return structure.PerformModuleWithPostgresDatabase(m, baseModule, baseModelName, baseTableName)
}

func PerfomBaseTemplate(m filesystem.Manager, goVersion, database string) error {
	prepareDirectoriesForScaffold(m)

	components := []filesystem.File{
		sharedComponent.MakeGoMod(m, goVersion),
		baseTemplateComponent.MakeEntry(m, baseModule, baseModelName),
		baseTemplateComponent.MakeConfig(m),
		baseTemplateComponent.MakeEnvironmentExample(m),
		baseTemplateComponent.MakeMakefile(m),
	}

	for _, c := range components {
		if err := m.GenerateFile(c); err != nil {
			return err
		}
	}

	if err := resource.PerformDockerContainer(m, goVersion); err != nil {
		return err
	}

	if err := resource.PerformDockerComposeWithPostgres(m); err != nil {
		return err
	}

	if err := resource.PerformGitignore(m); err != nil {
		return err
	}

	return structure.PerformModuleWithCustomDatabase(m, baseModule, baseModelName, database)
}

func prepareDirectoriesForScaffold(m filesystem.Manager) {
	m.GenerateNestedDirectories(
		m.MainDirectory(),
		[]string{"internal"},
	)
}
