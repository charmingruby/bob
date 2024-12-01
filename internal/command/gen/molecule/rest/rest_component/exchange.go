package rest_component

import (
	"fmt"

	"github.com/charmingruby/bob/internal/command/gen/atom"
	"github.com/charmingruby/bob/internal/command/gen/molecule/rest/structure"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/constant"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

func makeExchange(m component.Manager, module, name, exchange string) filesystem.File {
	component := atom.New(atom.ComponentInput{
		Module: module,
		Name:   name,
		Suffix: exchange,
		DestinationDirectory: m.AppendToModuleDirectory(
			module,
			fmt.Sprintf("transport/rest/dto/%s", exchange),
		),
		HasTest: false,
	})

	return atom.MakeCustomComponent(atom.CustomComponentInput{
		BaseComponent: *component,
		TemplateName:  constant.REST_EXCHANGE_TEMPLATE,
		TemplateData:  structure.NewExchangeData(exchange, name),
		FileName:      name,
		FileSuffix:    exchange,
	})
}

func MakeRequest(m component.Manager, module, handlerName string) filesystem.File {
	exchangeName := "request"

	return makeExchange(m, module, handlerName, exchangeName)
}

func MakeResponse(m component.Manager, module, handlerName string) filesystem.File {
	exchangeName := "response"

	return makeExchange(m, module, handlerName, exchangeName)
}
