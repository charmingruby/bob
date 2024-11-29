package atom

import (
	"github.com/charmingruby/bob/internal/command/gen/atom"
	"github.com/charmingruby/bob/internal/command/gen/molecule/rest/structure"
	"github.com/charmingruby/bob/internal/command/shared/fs"
)

func MakeRestRegistryComponent(destinationDirectory, sourcePath, module string) fs.File {
	return atom.MakeRegistryComponent(atom.RegistryParams{
		Module:       module,
		TemplateName: "rest_registry",
		TemplateData: structure.RestRegistryData{
			Module:     module,
			SourcePath: sourcePath,
		},
		RegistryName:         "endpoint",
		DestinationDirectory: destinationDirectory,
	})
}
