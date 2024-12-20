package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/organism/bootstrap/constant"
	"github.com/charmingruby/bob/internal/component/organism/bootstrap/data"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/shared"
)

func MakeContainer(m filesystem.Manager, goVersion string) filesystem.File {
	return base.New(base.ComponentInput{
		DestinationDirectory: m.RootDirectory,
	}).Componetize(
		shared.CREATE_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.CONTAINER_TEMPLATE,
			TemplateData: data.NewContainerData(goVersion),
			FileName:     "Dockerfile",
			Extension:    shared.NO_EXTENSION,
		})
}
