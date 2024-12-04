package rest_component

import (
	"github.com/charmingruby/bob/internal/command/gen/molecule/rest/structure"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/constant"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

func HandlerPath() string {
	return "transport/rest/endpoint"
}

func MakeHandlerComponent(m filesystem.Manager, module, name string) filesystem.File {
	return component.New(component.ComponentInput{
		Module:               module,
		Name:                 name,
		Suffix:               "handler",
		DestinationDirectory: filesystem.ModulePath(m.SourceDirectory, module, HandlerPath()),
	}).Componetize(component.ComponetizeInput{
		TemplateName: constant.REST_HANDLER_TEMPLATE,
		TemplateData: structure.NewHandlerData(
			m.DependencyPath(),
			constant.COMMON_MODULE,
			module,
			name,
		),
		FileName:   name,
		FileSuffix: "handler",
	})
}
