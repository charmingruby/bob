package atom

import (
	"github.com/charmingruby/bob/internal/command/generate/atom"
	"github.com/charmingruby/bob/internal/command/generate/molecule/rest/structure"
	"github.com/charmingruby/bob/internal/command/shared/fs"
)

func MakeRestRegistryComponent(destinationDirectory, sourceDirectory, module string) fs.File {
	return atom.MakeRegistryComponent(atom.RegistryParams{
		Module:       module,
		TemplateName: "rest_registry",
		TemplateData: structure.RestRegistryData{
			Module:          module,
			SourceDirectory: sourceDirectory,
		},
		RegistryName:         "endpoint",
		DestinationDirectory: destinationDirectory,
	})
}
