package component

import (
	"github.com/charmingruby/bob/internal/component/architecture/bundle/rest/data"
	"github.com/charmingruby/bob/internal/component/architecture/unit"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeHandlerRegistry(destinationDirectory, sourcePath, module string) filesystem.File {
	template := "architecture/bundle/rest/handler_registry"

	return unit.MakeRegistry(unit.RegistryParams{
		Package:      module,
		TemplateName: template,
		TemplateData: data.RestRegistryData{
			Module:     module,
			SourcePath: sourcePath,
		},
		RegistryName:         definition.HANDLER_PACKAGE,
		DestinationDirectory: destinationDirectory,
	})
}
