package atom

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/shared"
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
		shared.ADD_COMMAND,
		base.ComponetizeInput{
			TemplateName: params.TemplateName,
			TemplateData: params.TemplateData,
			FileName:     params.RegistryName,
			FileSuffix:   "",
		})
}
