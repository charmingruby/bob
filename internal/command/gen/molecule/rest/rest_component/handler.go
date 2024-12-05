package rest_component

import (
	restConstant "github.com/charmingruby/bob/internal/command/gen/molecule/rest/constant"
	"github.com/charmingruby/bob/internal/command/gen/molecule/rest/structure"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/constant"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

const handlerPath = "transport/rest/endpoint"

func MakeHandlerComponent(m filesystem.Manager, module, name string) filesystem.File {
	return component.New(component.ComponentInput{
		Module:               module,
		Name:                 name,
		Suffix:               "handler",
		DestinationDirectory: filesystem.ModulePath(m.SourceDirectory, module, handlerPath),
	}).Componetize(component.ComponetizeInput{
		TemplateName: restConstant.REST_HANDLER_TEMPLATE,
		TemplateData: structure.NewHandlerData(
			m.DependencyPath(),
			constant.SHARED_MODULE,
			module,
			name,
		),
		FileName:   name,
		FileSuffix: "handler",
	})
}
