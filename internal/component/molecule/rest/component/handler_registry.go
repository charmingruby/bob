package component

import (
	"github.com/charmingruby/bob/internal/component/atom"
	"github.com/charmingruby/bob/internal/component/molecule/constant"
	"github.com/charmingruby/bob/internal/component/molecule/rest/data"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/scaffold"
)

func MakeHandlerRegistry(destinationDirectory, sourcePath, module string) filesystem.File {
	return atom.MakeRegistry(atom.RegistryParams{
		Package:      module,
		TemplateName: constant.REST_HANDLER_REGISTRY_TEMPLATE,
		TemplateData: data.RestRegistryData{
			Module:     module,
			SourcePath: sourcePath,
		},
		RegistryName:         scaffold.HANDLER_PACKAGE,
		DestinationDirectory: destinationDirectory,
	})
}
