package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/organism/bootstrap/constant"
	"github.com/charmingruby/bob/internal/component/organism/bootstrap/data"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/shared"
)

func MakeGoMod(m filesystem.Manager, goVersion string) filesystem.File {
	return base.New(base.ComponentInput{
		DestinationDirectory: m.MainDirectory(),
	}).Componetize(
		shared.BOOTSTRAP_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.GO_MOD_TEMPLATE,
			TemplateData: data.NewGoModData(m.Data, goVersion),
			FileName:     "go",
			Extension:    "mod",
		})
}
