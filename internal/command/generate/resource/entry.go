package resource

import (
	"github.com/charmingruby/bob/internal/command/generate/brick"
	"github.com/charmingruby/bob/internal/command/shared/fs"
)

type entryBrickParams struct {
	Module               string
	EntryName            string
	DestinationDirectory string
	TemplateName         string
	TemplateData         any
}

func makeEntryBrick(params entryBrickParams) fs.File {
	component := brick.New(brick.ComponentInput{
		Module:    params.Module,
		Name:      params.EntryName,
		Suffix:    "",
		Directory: params.DestinationDirectory,
		HasTest:   false,
	})

	return brick.MakeCustomComponent(brick.CustomComponentInput{
		BaseComponent: *component,
		TemplateName:  params.TemplateName,
		TemplateData:  params.TemplateData,
		FileName:      params.EntryName,
		FileSuffix:    "",
	}, params.TemplateData)
}
