package unit

import (
	"fmt"

	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

type serviceData struct {
	ServiceName      string
	LowerCaseModel   string
	CapitalizedModel string
	SourcePath       string
	Module           string
}

func newServiceData(name, model, sourcePath, module string) serviceData {
	return serviceData{
		ServiceName:      base.CapitalizedFormat(name),
		LowerCaseModel:   base.LowerCaseFormat(model),
		CapitalizedModel: base.CapitalizedFormat(model),
		SourcePath:       sourcePath,
		Module:           base.SnakeCaseFormat(module),
	}
}

func MakeService(m filesystem.Manager, module, serviceName, modelToBeManaged string) filesystem.File {
	prepareDirectoriesForService(m, module)

	template := TemplatePath("service")

	destination := definition.CorePath(m.ModuleDirectory(module), []string{definition.SERVICE_PACKAGE})

	content := fmt.Sprintf("%s service", serviceName)

	return base.New(base.ComponentInput{
		Identifier:           base.BuildIdentifier(module, content, destination),
		Package:              module,
		Name:                 serviceName,
		Suffix:               "service",
		DestinationDirectory: destination,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			TemplateData: newServiceData(serviceName, modelToBeManaged, m.DependencyPath(), module),
			FileName:     serviceName,
			FileSuffix:   "service",
		})
}

func prepareDirectoriesForService(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{definition.CORE_PACKAGE, definition.SERVICE_PACKAGE},
	)
}
