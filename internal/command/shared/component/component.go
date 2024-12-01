package component

import (
	"github.com/charmingruby/bob/internal/command/shared/component/constant"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

type Component struct {
	DestinationDirectory string
	Module               string
	Name                 string
	Suffix               string
	HasTest              bool
}

type ComponentInput struct {
	Module               string
	Name                 string
	Suffix               string
	DestinationDirectory string
	HasTest              bool
}

func New(in ComponentInput) *Component {
	component := &Component{
		Module:               in.Module,
		Name:                 in.Name,
		DestinationDirectory: in.DestinationDirectory,
		Suffix:               in.Suffix,
		HasTest:              in.HasTest,
	}

	return component
}

type ComponetizeInput struct {
	TemplateName string
	TemplateData any
	FileName     string
	FileSuffix   string
}

func (c *Component) Componetize(in ComponetizeInput) filesystem.File {
	return filesystem.File{
		CommandType:          constant.GENERATE_COMMAND,
		TemplateName:         in.TemplateName,
		TemplateData:         in.TemplateData,
		FileName:             in.FileName,
		FileSuffix:           in.FileSuffix,
		DestinationDirectory: c.DestinationDirectory,
		HasTest:              c.HasTest,
	}
}
