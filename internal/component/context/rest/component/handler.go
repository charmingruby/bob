package component

import (
	"fmt"

	"github.com/charmingruby/bob/internal/component/context/rest"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/charmingruby/bob/pkg/formatter"
)

type handlerData struct {
	SourcePath  string
	Module      string
	Name        string
	ServiceName string
}

func newHandlerData(sourcePath, module, name string) handlerData {
	return handlerData{
		SourcePath:  sourcePath,
		Module:      module,
		Name:        formatter.ToCamelCase(name),
		ServiceName: base.PublicNameFormat(name),
	}
}

func MakeHandler(m filesystem.Manager, module, name string) filesystem.File {
	prepareDirectoriesForHandler(m, module)

	template := rest.TemplatePath("component/handler")

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
			TemplateData: newHandlerData(
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
