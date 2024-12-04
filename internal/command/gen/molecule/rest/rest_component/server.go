package rest_component

import (
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/constant"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

func ServerPath() string {
	return "transport/rest"
}

func MakeServerComponent(m filesystem.Manager) filesystem.File {
	return component.New(component.ComponentInput{
		Module:               "rest",
		DestinationDirectory: m.AppendToModuleDirectory(constant.COMMON_MODULE, ServerPath()),
	}).Componetize(component.ComponetizeInput{
		TemplateName: constant.REST_SERVER,
		FileName:     "server",
	})
}
