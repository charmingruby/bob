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

func prepareDirectoriesForExchange(m filesystem.Manager, module, exchange string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{scaffold.TRANSPORT_PACKAGE, scaffold.REST_PACKAGE, scaffold.DTO_PACKAGE, exchange},
	)
}

func MakeRequest(m filesystem.Manager, module, handlerName string) filesystem.File {
	exchange := "request"
	prepareDirectoriesForExchange(m, module, exchange)
	return makeExchangeComponent(m, module, handlerName, exchange)
}

func MakeResponse(m filesystem.Manager, module, handlerName string) filesystem.File {
	exchange := "response"
	prepareDirectoriesForExchange(m, module, exchange)
	return makeExchangeComponent(m, module, handlerName, exchange)
}
