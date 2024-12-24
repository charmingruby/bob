package component

import (
	"fmt"

	"github.com/charmingruby/bob/internal/component/architecture/bundle/rest/data"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeHandler(m filesystem.Manager, module, name string) filesystem.File {
	prepareDirectoriesForHandler(m, module)

	template := "architecture/bundle/rest/handler"

	destination := definition.TransportPath(
		m.ModuleDirectory(module),
		definition.REST_PACKAGE,
		[]string{definition.HANDLER_PACKAGE},
	)

	content := fmt.Sprintf("%s endpoint handler", name)

	return base.New(base.ComponentInput{
		Identifier:           base.BuildIdentifier(module, content, destination),
		Package:              module,
		Name:                 name,
		Suffix:               "handler",
		DestinationDirectory: destination,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			TemplateData: data.NewHandlerData(
				m.DependencyPath(),
				module,
				name,
			),
			FileName:   name,
			FileSuffix: "handler",
		})
}

func prepareDirectoriesForHandler(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{definition.TRANSPORT_PACKAGE, definition.REST_PACKAGE, definition.HANDLER_PACKAGE},
	)
}
