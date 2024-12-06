package rest_component

import (
	"github.com/charmingruby/bob/internal/command/gen/molecule/rest/constant"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
	"github.com/charmingruby/bob/internal/command/shared/scaffold"
)

func MakeBaseServerMiddlewareComponent(m filesystem.Manager) filesystem.File {
	return component.New(component.ComponentInput{
		Package: scaffold.SERVICE_PACKAGE,
		DestinationDirectory: scaffold.TransportPath(
			m.ModuleDirectory(scaffold.SHARED_MODULE),
			scaffold.REST_PACKAGE,
			nil,
		),
	}).Componetize(component.ComponetizeInput{
		TemplateName: constant.REST_BASE_SERVER_MIDDLEWARE,
		FileName:     "middleware",
	})
}
