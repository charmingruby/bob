package atom

import (
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

type RegistryParams struct {
	Module               string
	RegistryName         string
	DestinationDirectory string
	TemplateName         string
	TemplateData         any
}

func MakeRegistryComponent(params RegistryParams) filesystem.File {
	component := New(ComponentInput{
		Module:               params.Module,
		Name:                 params.RegistryName,
		Suffix:               "",
		DestinationDirectory: params.DestinationDirectory,
		HasTest:              false,
	})

	return MakeCustomComponent(CustomComponentInput{
		BaseComponent: *component,
		TemplateName:  params.TemplateName,
		TemplateData:  params.TemplateData,
		FileName:      params.RegistryName,
		FileSuffix:    "",
	})
}
