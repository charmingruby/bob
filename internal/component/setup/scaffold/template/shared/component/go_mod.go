package component

import (
	"github.com/charmingruby/bob/internal/component/setup/scaffold/template/shared/data"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeGoMod(m filesystem.Manager, goVersion string) filesystem.File {
	template := "setup/scaffold/go_mod"

	destination := m.RootDirectory

	resource := "go"

	content := "go mod"

	return base.New(base.ComponentInput{
		Identifier:           base.BuildNonModuleIdentifier(resource, content, destination),
		DestinationDirectory: destination,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			TemplateData: data.NewGoModData(m.Data, goVersion),
			FileName:     "go",
			Extension:    "mod",
		})
}
