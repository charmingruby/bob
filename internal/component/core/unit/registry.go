package unit

import (
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

type RegistryParams struct {
	Identifier           string
	Package              string
	RegistryName         string
	DestinationDirectory string
	TemplateName         string
	TemplateData         any
}

func MakeRegistry(params RegistryParams) filesystem.File {
	return base.New(base.ComponentInput{
		Identifier:           params.Identifier,
		Package:              params.Package,
		Name:                 params.RegistryName,
		DestinationDirectory: params.DestinationDirectory,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: params.TemplateName,
			TemplateData: params.TemplateData,
			FileName:     params.RegistryName,
			FileSuffix:   "",
		})
}
