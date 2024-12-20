package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/organism/bootstrap/constant"
	"github.com/charmingruby/bob/internal/component/organism/bootstrap/data"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeContainer(m filesystem.Manager, goVersion string) filesystem.File {
	return base.New(base.ComponentInput{
		DestinationDirectory: m.RootDirectory,
	}).Componetize(
		definition.CREATE_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.CONTAINER_TEMPLATE,
			TemplateData: data.NewContainerData(goVersion),
			FileName:     "Dockerfile",
			Extension:    definition.NO_EXTENSION,
		})
}
