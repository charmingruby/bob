package atom

import (
	"github.com/charmingruby/bob/internal/command/shared/component/constant"
	"github.com/charmingruby/bob/internal/command/shared/component/structure"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

type Component struct {
	DestinationDirectory string
	Module               string
	Name                 string
	Suffix               string
	Data                 any
	HasTest              bool
}

type ComponentInput struct {
	Module               string
	Name                 string
	Suffix               string
	DestinationDirectory string
	HasTest              bool
}

type ComponentOption func(*Component)

func New(in ComponentInput, opts ...ComponentOption) *Component {
	component := &Component{
		Module:               in.Module,
		Name:                 in.Name,
		DestinationDirectory: in.DestinationDirectory,
		Suffix:               in.Suffix,
		HasTest:              in.HasTest,
	}

	for _, opt := range opts {
		opt(component)
	}

	return component
}

func WithDefaultTemplate() ComponentOption {
	return func(s *Component) {
		s.Data = structure.NewPureData(s.Name)
	}
}

func WithModuleDependenciesTemplate(sourcePath string) ComponentOption {
	return func(s *Component) {
		s.Data = structure.NewDependentPackageData(sourcePath, s.Module, s.Name)
	}
}

func NewFileFromAtom(component Component, template string) filesystem.File {
	return filesystem.File{
		CommandType:          constant.GENERATE_COMMAND,
		TemplateName:         template,
		TemplateData:         component.Data,
		FileName:             component.Name,
		FileSuffix:           component.Suffix,
		DestinationDirectory: component.DestinationDirectory,
		HasTest:              component.HasTest,
	}
}
