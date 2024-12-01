package atom

import (
	"github.com/charmingruby/bob/internal/command/shared/component"
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
	return component.New(component.ComponentInput{
		Module:               params.Module,
		Name:                 params.RegistryName,
		DestinationDirectory: params.DestinationDirectory,
	}).Componetize(component.ComponetizeInput{
		TemplateName: params.TemplateName,
		TemplateData: params.TemplateData,
		FileName:     params.RegistryName,
		FileSuffix:   "",
	})
}
