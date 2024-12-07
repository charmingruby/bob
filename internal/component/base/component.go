package base

import (
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/scaffold"
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

func (c *Component) Componetize(in ComponetizeInput) filesystem.File {
	return filesystem.File{
		CommandType:          scaffold.GENERATE_COMMAND,
		TemplateName:         in.TemplateName,
		TemplateData:         in.TemplateData,
		FileName:             in.FileName,
		FileSuffix:           in.FileSuffix,
		DestinationDirectory: c.DestinationDirectory,
		HasTest:              c.HasTest,
	}
}
