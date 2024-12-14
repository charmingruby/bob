package base

import (
	"github.com/charmingruby/bob/internal/filesystem"
)

type Component struct {
	Package              string
	Name                 string
	Suffix               string
	DestinationDirectory string
	HasTest              bool
}

type ComponentInput struct {
	Package              string
	Name                 string
	Suffix               string
	DestinationDirectory string
	HasTest              bool
}

func New(in ComponentInput) *Component {
	component := &Component{
		Package:              in.Package,
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

func (c *Component) Componetize(commandType string, in ComponetizeInput) filesystem.File {
	return filesystem.File{
		CommandType:          commandType,
		TemplateName:         in.TemplateName,
		TemplateData:         in.TemplateData,
		FileName:             in.FileName,
		FileSuffix:           in.FileSuffix,
		DestinationDirectory: c.DestinationDirectory,
		HasTest:              c.HasTest,
	}
}
