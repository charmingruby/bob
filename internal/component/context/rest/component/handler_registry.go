package component

import (
	"github.com/charmingruby/bob/internal/component/core/unit"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

type restRegistryData struct {
	Module     string
	SourcePath string
}

func newRestRegistryData(module, sourcePath string) restRegistryData {
	return restRegistryData{
		Module:     base.ModuleFormat(module),
		SourcePath: sourcePath,
	}
}

func MakeHandlerRegistry(destinationDirectory, sourcePath, module string) filesystem.File {
	template := "architecture/bundle/rest/handler_registry"

	content := "endpoint entrypoint"

	return unit.MakeRegistry(unit.RegistryParams{
		Identifier:           base.BuildIdentifier(module, content, destinationDirectory),
		Package:              module,
		TemplateName:         template,
		TemplateData:         newRestRegistryData(module, sourcePath),
		RegistryName:         definition.HANDLER_PACKAGE,
		DestinationDirectory: destinationDirectory,
	})
}
