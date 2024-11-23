package single

import (
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/formatter"
)

type Single struct {
	Directory string
	Module    string
	Package   component.Package
	Name      string
	Suffix    string
	Data      any
	HasTest   bool
}

type SingleInput struct {
	Module      string
	Name        string
	Suffix      string
	PackageName string
	HasTest     bool
}

type SingleOption func(*Single)

func New(in SingleInput, opts ...SingleOption) *Single {
	component := &Single{
		Module: in.Module,
		Name:   in.Name,
		Package: component.Package{
			Name: in.PackageName,
		},
		Suffix:  in.Suffix,
		HasTest: in.HasTest,
	}

	component.format()

	for _, opt := range opts {
		opt(component)
	}

	return component
}

type SingleDefaultTemplate struct {
	Package                   string
	PackageRegistry           string
	PackageRegistryIdentifier string
	Name                      string
}

func WithDefaultTemplate() SingleOption {
	return func(s *Single) {
		var fullName = s.Name
		if s.Suffix != "" {
			fullName += s.Suffix
		}

		s.Data = SingleDefaultTemplate{
			Package:                   s.Package.Name,
			PackageRegistry:           s.Package.Registry,
			PackageRegistryIdentifier: s.Package.RegistryIdentifier,
			Name:                      fullName,
		}
	}
}

func (r *Single) format() {
	r.Module = formatter.ToSnakeCase(r.Module)
	r.Name = formatter.ToCamelCase(r.Name)

	if r.Suffix != "" {
		r.Suffix = formatter.ToCamelCase(r.Suffix)
	}

	r.Package.Format()
}
