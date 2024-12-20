package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/molecule/rest/constant"
	"github.com/charmingruby/bob/internal/component/molecule/rest/data"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/shared"
)

func makeExchange(m filesystem.Manager, module, name, exchange string) filesystem.File {
	return base.New(base.ComponentInput{
		Package: module,
		Name:    name,
		Suffix:  exchange,
		DestinationDirectory: shared.TransportPath(
			m.ModuleDirectory(module),
			shared.REST_PACKAGE,
			[]string{shared.DTO_PACKAGE, exchange},
		),
	}).Componetize(
		shared.ADD_COMMAND,
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
		[]string{shared.TRANSPORT_PACKAGE, shared.REST_PACKAGE, shared.DTO_PACKAGE, exchange},
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
