package single

import (
	"strings"

	"github.com/charmingruby/bob/internal/command/shared/formatter"
)

type Single struct {
	Identifier string
	Directory  string
	Module     string
	Name       string
	FullName   string
	Suffix     string
	Package    Package
	Data       any
	ActionType string
	HasTest    bool
}

type Package struct {
	Name               string
	RegistryIdentifier string
	Registry           string
}

type SingleInput struct {
	Identifier    string
	Module        string
	Name          string
	Suffix        string
	PackageName   string
	ActionType    string
	HasTest       bool
	HasSuffixName bool
}

type SingleOption func(*Single)

func New(in SingleInput, opts ...SingleOption) *Single {
	component := &Single{
		Identifier: in.Identifier,
		Module:     in.Module,
		Name:       in.Name,
		Package: Package{
			Name: in.PackageName,
		},
		Suffix:     in.Suffix,
		ActionType: in.ActionType,
		HasTest:    in.HasTest,
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
	return func(c *Single) {
		c.Data = SingleDefaultTemplate{
			Package:                   c.Package.Name,
			PackageRegistry:           c.Package.Registry,
			PackageRegistryIdentifier: c.Package.RegistryIdentifier,
			Name:                      c.FullName,
		}
	}
}

func (r *Single) format() {
	r.Module = formatter.ToSnakeCase(r.Module)
	r.Name = formatter.ToCamelCase(r.Name)
	r.FullName = r.Name
	r.Package.Name = formatter.ToSnakeCase(r.Package.Name)
	r.Package.Registry = formatter.ToCamelCase(r.Package.Name)
	r.Package.RegistryIdentifier = formatter.ToLowerCase(string(r.Package.Name[0]))
	if r.Suffix != "" {
		r.FullName = formatter.ToCamelCase(r.Name + strings.ToTitle(r.Suffix))
		r.Suffix = "_" + formatter.ToSnakeCase(r.Suffix)
	}
}
