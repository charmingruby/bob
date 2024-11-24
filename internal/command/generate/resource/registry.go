package resource

import (
	"github.com/charmingruby/bob/internal/command/generate/brick"
	"github.com/charmingruby/bob/internal/command/shared/fs"
)

type registryBrickParams struct {
	Module               string
	RegistryName         string
	DestinationDirectory string
	TemplateName         string
	TemplateData         any
}

func makeRegistryBrick(params registryBrickParams) fs.File {
	component := brick.New(brick.ComponentInput{
		Module:    params.Module,
		Name:      params.RegistryName,
		Suffix:    "",
		Directory: params.DestinationDirectory,
		HasTest:   false,
	})

	return brick.MakeCustomComponent(brick.CustomComponentInput{
		BaseComponent: *component,
		TemplateName:  params.TemplateName,
		TemplateData:  params.TemplateData,
		FileName:      params.RegistryName,
		FileSuffix:    "",
	}, params.TemplateData)
}
