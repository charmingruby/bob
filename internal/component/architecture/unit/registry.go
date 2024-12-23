package unit

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

type RegistryParams struct {
	Package              string
	RegistryName         string
	DestinationDirectory string
	TemplateName         string
	TemplateData         any
}

func MakeRegistry(params RegistryParams) filesystem.File {
	return base.New(base.ComponentInput{
		Package:              params.Package,
		Name:                 params.RegistryName,
		DestinationDirectory: params.DestinationDirectory,
	}).Componetize(
		definition.ADD_COMMAND,
		base.ComponetizeInput{
			TemplateName: params.TemplateName,
			TemplateData: params.TemplateData,
			FileName:     params.RegistryName,
			FileSuffix:   "",
		})
}
