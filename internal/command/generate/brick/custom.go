package brick

import (
	"github.com/charmingruby/bob/internal/command/shared/constant"
	"github.com/charmingruby/bob/internal/command/shared/fs"
)

type CustomComponentInput struct {
	BaseComponent Component
	TemplateName  string
	TemplateData  any
	FileName      string
	FileSuffix    string
}

func MakeCustomComponent(in CustomComponentInput, data any) fs.File {
	file := fs.File{
		CommandType:          constant.GENERATE_COMMAND,
		TemplateName:         in.TemplateName,
		TemplateData:         in.TemplateData,
		FileName:             in.FileName,
		FileSuffix:           in.FileSuffix,
		DestinationDirectory: in.BaseComponent.Directory,
		HasTest:              in.BaseComponent.HasTest,
	}

	return file
}
