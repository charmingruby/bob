package rest_component

import (
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/constant"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

func BaseServerMiddlewarePath() string {
	return "transport/rest"
}

func MakeBaseServerMiddlewareComponent(m filesystem.Manager) filesystem.File {
	return component.New(component.ComponentInput{
		Module:               "rest",
		DestinationDirectory: m.AppendToModuleDirectory(constant.COMMON_MODULE, BaseServerMiddlewarePath()),
	}).Componetize(component.ComponetizeInput{
		TemplateName: constant.REST_BASE_SERVER_MIDDLEWARE,
		FileName:     "middleware",
	})
}
