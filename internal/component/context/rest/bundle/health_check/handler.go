package health_check

import (
	"github.com/charmingruby/bob/internal/component/context/rest"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

type HandlerData struct {
	SourcePath string
}

func newHandlerData(sourcePath string) HandlerData {
	return HandlerData{
		SourcePath: sourcePath,
	}
}

func MakeHandler(m filesystem.Manager) filesystem.File {
	template := rest.TemplatePath("bundle/health_check/handler")

	destination := definition.TransportPath(
		m.ModuleDirectory(definition.SHARED_MODULE),
		definition.REST_PACKAGE,
		[]string{definition.HANDLER_PACKAGE},
	)

	content := "health check handler"

	return base.New(base.ComponentInput{
		Identifier:           base.BuildIdentifier(definition.SHARED_MODULE, content, destination),
		DestinationDirectory: destination,
	}).Componetize(base.ComponetizeInput{
		TemplateName: template,
		TemplateData: newHandlerData(m.DependencyPath()),
		FileName:     "health_check",
		FileSuffix:   "handler",
	})
}
