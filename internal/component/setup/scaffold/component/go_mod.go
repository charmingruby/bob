package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/setup/scaffold/constant"
	"github.com/charmingruby/bob/internal/component/setup/scaffold/data"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeGoMod(m filesystem.Manager, goVersion string) filesystem.File {
	return base.New(base.ComponentInput{
		DestinationDirectory: m.MainDirectory(),
	}).Componetize(
		definition.CREATE_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.GO_MOD_TEMPLATE,
			TemplateData: data.NewGoModData(m.Data, goVersion),
			FileName:     "go",
			Extension:    "mod",
		})
}
