package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/setup/scaffold/template/shared/data"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeGoMod(m filesystem.Manager, goVersion string) filesystem.File {
	template := "setup/scaffold/go_mod"

	return base.New(base.ComponentInput{
		DestinationDirectory: m.MainDirectory(),
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			TemplateData: data.NewGoModData(m.Data, goVersion),
			FileName:     "go",
			Extension:    "mod",
		})
}
