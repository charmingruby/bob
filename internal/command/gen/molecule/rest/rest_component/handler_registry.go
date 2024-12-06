package rest_component

import (
	"github.com/charmingruby/bob/internal/command/gen/atom"
	"github.com/charmingruby/bob/internal/command/gen/molecule/rest/constant"
	"github.com/charmingruby/bob/internal/command/gen/molecule/rest/structure"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
	"github.com/charmingruby/bob/internal/command/shared/scaffold"
)

func MakeHandlerRegistryComponent(destinationDirectory, sourcePath, module string) filesystem.File {
	return atom.MakeRegistryComponent(atom.RegistryParams{
		Package:      module,
		TemplateName: constant.REST_HANDLER_REGISTRY_TEMPLATE,
		TemplateData: structure.RestRegistryData{
			Module:     module,
			SourcePath: sourcePath,
		},
		RegistryName:         scaffold.HANDLER_PACKAGE,
		DestinationDirectory: destinationDirectory,
	})
}
