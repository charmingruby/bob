package rest_component

import (
	"fmt"

	"github.com/charmingruby/bob/internal/command/gen/molecule/rest/structure"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/constant"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

func makeExchangeComponent(m filesystem.Manager, module, name, exchange string) filesystem.File {
	return component.New(component.ComponentInput{
		Module: module,
		Name:   name,
		Suffix: exchange,
		DestinationDirectory: m.AppendToModuleDirectory(
			module,
			fmt.Sprintf("transport/rest/dto/%s", exchange),
		),
	}).Componetize(component.ComponetizeInput{
		TemplateName: constant.REST_EXCHANGE_TEMPLATE,
		TemplateData: structure.NewExchangeData(exchange, name),
		FileName:     name,
		FileSuffix:   exchange,
	})
}

func MakeRequest(m filesystem.Manager, module, handlerName string) filesystem.File {
	exchangeName := "request"

	return makeExchangeComponent(m, module, handlerName, exchangeName)
}

func MakeResponse(m filesystem.Manager, module, handlerName string) filesystem.File {
	exchangeName := "response"

	return makeExchangeComponent(m, module, handlerName, exchangeName)
}
