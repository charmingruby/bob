package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/molecule/constant"
	"github.com/charmingruby/bob/internal/component/molecule/rest/data"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/scaffold"
)

func makeExchange(m filesystem.Manager, module, name, exchange string) filesystem.File {
	return base.New(base.ComponentInput{
		Package: module,
		Name:    name,
		Suffix:  exchange,
		DestinationDirectory: scaffold.TransportPath(
			m.ModuleDirectory(module),
			scaffold.REST_PACKAGE,
			[]string{scaffold.DTO_PACKAGE, exchange},
		),
	}).Componetize(base.ComponetizeInput{
		TemplateName: constant.REST_EXCHANGE_TEMPLATE,
		TemplateData: data.NewExchangeData(exchange, name),
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
	return makeExchange(m, module, handlerName, exchange)
}

func MakeResponse(m filesystem.Manager, module, handlerName string) filesystem.File {
	exchange := "response"
	prepareDirectoriesForExchange(m, module, exchange)
	return makeExchange(m, module, handlerName, exchange)
}
