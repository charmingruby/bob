package component

import (
	"fmt"

	"github.com/charmingruby/bob/internal/component/context/rest"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

type exchangeData struct {
	ActionName string
}

func newExchangeData(name string) exchangeData {
	return exchangeData{
		ActionName: base.CapitalizedFormat(name),
	}
}

func makeExchange(m filesystem.Manager, module, name, exchange string) filesystem.File {
	template := rest.TemplatePath("component/" + exchange)

	destination := definition.TransportPath(
		m.ModuleDirectory(module),
		definition.REST_PACKAGE,
		[]string{definition.DTO_PACKAGE, exchange},
	)

	content := fmt.Sprintf("%s %s", name, exchange)

	return base.New(base.ComponentInput{
		Identifier:           base.BuildIdentifier(module, content, destination),
		Package:              module,
		Name:                 name,
		Suffix:               exchange,
		DestinationDirectory: destination,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			TemplateData: newExchangeData(name),
			FileName:     name,
			FileSuffix:   exchange,
			Extension:    definition.GO_EXTENSION,
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
