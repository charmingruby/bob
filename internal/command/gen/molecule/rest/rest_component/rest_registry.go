package rest_component

import (
	"github.com/charmingruby/bob/internal/command/gen/atom"
	"github.com/charmingruby/bob/internal/command/gen/molecule/rest/constant"
	"github.com/charmingruby/bob/internal/command/gen/molecule/rest/structure"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

func MakeRestRegistryComponent(destinationDirectory, sourcePath, module string) filesystem.File {
	return atom.MakeRegistryComponent(atom.RegistryParams{
		Module:       module,
		TemplateName: constant.REST_REGISTRY_TEMPLATE,
		TemplateData: structure.RestRegistryData{
			Module:     module,
			SourcePath: sourcePath,
		},
		RegistryName:         "endpoint",
		DestinationDirectory: destinationDirectory,
	})
}
