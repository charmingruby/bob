package rest_component

import (
	restConstant "github.com/charmingruby/bob/internal/command/gen/molecule/rest/constant"
	"github.com/charmingruby/bob/internal/command/gen/molecule/rest/structure"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
	"github.com/charmingruby/bob/internal/command/shared/scaffold"
)

func makeExchangeComponent(m filesystem.Manager, module, name, exchange string) filesystem.File {
	return component.New(component.ComponentInput{
		Package: module,
		Name:    name,
		Suffix:  exchange,
		DestinationDirectory: scaffold.TransportPath(
			m.ModuleDirectory(module),
			scaffold.REST_PACKAGE,
			[]string{scaffold.DTO_PACKAGE, exchange},
		),
	}).Componetize(component.ComponetizeInput{
		TemplateName: restConstant.REST_EXCHANGE_TEMPLATE,
		TemplateData: structure.NewExchangeData(exchange, name),
		FileName:     name,
		FileSuffix:   exchange,
	})
}

func MakeRequest(m filesystem.Manager, module, handlerName string) filesystem.File {
	return makeExchangeComponent(m, module, handlerName, "request")
}

func MakeResponse(m filesystem.Manager, module, handlerName string) filesystem.File {
	return makeExchangeComponent(m, module, handlerName, "response")
}
