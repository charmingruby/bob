package base

import (
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/charmingruby/bob/pkg/util"
)

type Component struct {
	Identifier           string
	Package              string
	Name                 string
	Suffix               string
	DestinationDirectory string
	HasTest              bool
}

type ComponentInput struct {
	Identifier           string
	Package              string
	Name                 string
	Suffix               string
	DestinationDirectory string
	HasTest              bool
}

func New(in ComponentInput) *Component {
	component := &Component{
		Identifier:           in.Identifier,
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
	Extension    string
}

func (c *Component) Componetize(in ComponetizeInput) filesystem.File {
	var extension string = util.Ternary[string](in.Extension == "", definition.GO_EXTENSION, in.Extension)

	return filesystem.File{
		Identifier:           c.Identifier,
		TemplateName:         in.TemplateName,
		TemplateData:         in.TemplateData,
		FileName:             in.FileName,
		FileSuffix:           in.FileSuffix,
		DestinationDirectory: c.DestinationDirectory,
		HasTest:              c.HasTest,
		Extension:            extension,
	}
}
