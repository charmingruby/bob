package atom

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/filesystem"
)

type RegistryParams struct {
	Package              string
	RegistryName         string
	DestinationDirectory string
	TemplateName         string
	TemplateData         any
}

func MakeRegistryComponent(params RegistryParams) filesystem.File {
	return base.New(base.ComponentInput{
		Package:              params.Package,
		Name:                 params.RegistryName,
		DestinationDirectory: params.DestinationDirectory,
	}).Componetize(base.ComponetizeInput{
		TemplateName: params.TemplateName,
		TemplateData: params.TemplateData,
		FileName:     params.RegistryName,
		FileSuffix:   "",
	})
}
