package atom

import (
	"github.com/charmingruby/bob/internal/command/shared/fs"
)

type RegistryParams struct {
	Module               string
	RegistryName         string
	DestinationDirectory string
	TemplateName         string
	TemplateData         any
}

func MakeRegistryComponent(params RegistryParams) fs.File {
	component := New(ComponentInput{
		Module:    params.Module,
		Name:      params.RegistryName,
		Suffix:    "",
		Directory: params.DestinationDirectory,
		HasTest:   false,
	})

	return MakeCustomComponent(CustomComponentInput{
		BaseComponent: *component,
		TemplateName:  params.TemplateName,
		TemplateData:  params.TemplateData,
		FileName:      params.RegistryName,
		FileSuffix:    "",
	}, params.TemplateData)
}
