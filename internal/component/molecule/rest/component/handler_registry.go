package component

import (
	"github.com/charmingruby/bob/internal/component/atom"
	"github.com/charmingruby/bob/internal/component/molecule/rest/constant"
	"github.com/charmingruby/bob/internal/component/molecule/rest/data"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeHandlerRegistry(destinationDirectory, sourcePath, module string) filesystem.File {
	return atom.MakeRegistry(atom.RegistryParams{
		Package:      module,
		TemplateName: constant.REST_HANDLER_REGISTRY_TEMPLATE,
		TemplateData: data.RestRegistryData{
			Module:     module,
			SourcePath: sourcePath,
		},
		RegistryName:         definition.HANDLER_PACKAGE,
		DestinationDirectory: destinationDirectory,
	})
}
