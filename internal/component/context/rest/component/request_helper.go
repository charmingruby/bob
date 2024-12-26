package component

import (
	"github.com/charmingruby/bob/internal/component/context/rest"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

type requestHelperData struct {
	SourcePath string
}

func newRequestHelperData(sourcePath string) requestHelperData {
	return requestHelperData{
		SourcePath: sourcePath,
	}
}

func MakeRequestHelper(m filesystem.Manager) filesystem.File {
	prepareDirectoriesForRequestHelper(m, definition.SHARED_MODULE)

	template := rest.TemplatePath("component/request_helper")

	destination := definition.TransportPath(
		m.ModuleDirectory(definition.SHARED_MODULE),
		definition.REST_PACKAGE,
		nil,
	)

	content := "rest request helper"

	return base.New(base.ComponentInput{
		Identifier:           base.BuildIdentifier(definition.SHARED_MODULE, content, destination),
		Package:              definition.REST_PACKAGE,
		DestinationDirectory: destination,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			TemplateData: newRequestHelperData(m.DependencyPath()),
			FileName:     "request",
		})
}

func prepareDirectoriesForRequestHelper(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{module, definition.TRANSPORT_PACKAGE, definition.REST_PACKAGE},
	)
}
