package rest_component

import (
	restConstant "github.com/charmingruby/bob/internal/command/gen/molecule/rest/constant"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/constant"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

const middlewarePath = "transport/rest"

func MakeBaseServerMiddlewareComponent(m filesystem.Manager) filesystem.File {
	return component.New(component.ComponentInput{
		Module:               "rest",
		DestinationDirectory: m.AppendToModuleDirectory(constant.SHARED_MODULE, middlewarePath),
	}).Componetize(component.ComponetizeInput{
		TemplateName: restConstant.REST_BASE_SERVER_MIDDLEWARE,
		FileName:     "middleware",
	})
}
