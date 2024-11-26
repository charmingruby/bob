package atom

import (
	"github.com/charmingruby/bob/internal/command/shared/component/constant"
	"github.com/charmingruby/bob/internal/command/shared/component/structure"
	"github.com/charmingruby/bob/internal/command/shared/fs"
)

type Component struct {
	Directory string
	Module    string
	Name      string
	Suffix    string
	Data      any
	HasTest   bool
}

type ComponentInput struct {
	Module    string
	Name      string
	Suffix    string
	Directory string
	HasTest   bool
}

type ComponentOption func(*Component)

func New(in ComponentInput, opts ...ComponentOption) *Component {
	component := &Component{
		Module:    in.Module,
		Name:      in.Name,
		Directory: in.Directory,
		Suffix:    in.Suffix,
		HasTest:   in.HasTest,
	}

	for _, opt := range opts {
		opt(component)
	}

	return component
}

func WithDefaultTemplate() ComponentOption {
	return func(s *Component) {
		s.Data = structure.Pure{
			Name: s.Name,
		}
	}
}

func WithModuleDependenciesTemplate(sourcePath string) ComponentOption {
	return func(s *Component) {
		s.Data = structure.DependentPackage{
			Module:     s.Module,
			SourcePath: sourcePath,
			Name:       s.Name,
		}
	}
}

func NewFileFromAtom(component Component, template string) fs.File {
	return fs.File{
		CommandType:          constant.GENERATE_COMMAND,
		TemplateName:         template,
		TemplateData:         component.Data,
		FileName:             component.Name,
		FileSuffix:           component.Suffix,
		DestinationDirectory: component.Directory,
		HasTest:              component.HasTest,
	}
}
