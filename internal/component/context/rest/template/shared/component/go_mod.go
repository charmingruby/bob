package component

import (
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

type goModData struct {
	RepoURL   string
	GoVersion string
}

func newGoModData(
	repoURL, goVersion string,
) goModData {
	return goModData{
		RepoURL:   repoURL,
		GoVersion: goVersion,
	}
}

func MakeGoMod(m filesystem.Manager, goVersion string) filesystem.File {
	template := "shared/template/go_mod"

	destination := m.MainDirectory()

	resource := "go"

	content := "go mod"

	return base.New(base.ComponentInput{
		Identifier:           base.BuildNonModuleIdentifier(resource, content, destination),
		DestinationDirectory: destination,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			TemplateData: newGoModData(m.Data, goVersion),
			FileName:     "go",
			Extension:    "mod",
		})
}
