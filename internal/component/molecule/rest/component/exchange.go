package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/molecule/rest/constant"
	"github.com/charmingruby/bob/internal/component/molecule/rest/data"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func makeExchange(m filesystem.Manager, module, name, exchange string) filesystem.File {
	return base.New(base.ComponentInput{
		Package: module,
		Name:    name,
		Suffix:  exchange,
		DestinationDirectory: definition.TransportPath(
			m.ModuleDirectory(module),
			definition.REST_PACKAGE,
			[]string{definition.DTO_PACKAGE, exchange},
		),
	}).Componetize(
		definition.ADD_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.REST_EXCHANGE_TEMPLATE,
			TemplateData: data.NewExchangeData(exchange, name),
			FileName:     name,
			FileSuffix:   exchange,
		})
}

func prepareDirectoriesForExchange(m filesystem.Manager, module, exchange string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{definition.TRANSPORT_PACKAGE, definition.REST_PACKAGE, definition.DTO_PACKAGE, exchange},
	)
}

func MakeRequest(m filesystem.Manager, module, handlerName string) filesystem.File {
	exchange := "request"
	prepareDirectoriesForExchange(m, module, exchange)
	return makeExchange(m, module, handlerName, exchange)
}

func MakeResponse(m filesystem.Manager, module, handlerName string) filesystem.File {
	exchange := "response"
	prepareDirectoriesForExchange(m, module, exchange)
	return makeExchange(m, module, handlerName, exchange)
}
