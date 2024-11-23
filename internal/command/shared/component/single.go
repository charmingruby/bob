package component

import "strings"

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

func WithDefaultTemplateParams() SingleOption {
	return func(c *Single) {
		c.Data = DefaultTemplateParams{
			Package:                   c.Package.Name,
			PackageRegistry:           c.Package.Registry,
			PackageRegistryIdentifier: c.Package.RegistryIdentifier,
			Name:                      c.FullName,
		}
	}
}

func (r *Single) format() {
	r.Module = toSnakeCase(r.Module)
	r.Name = toCamelCase(r.Name)
	r.FullName = r.Name
	r.Package.Name = toSnakeCase(r.Package.Name)
	r.Package.Registry = toCamelCase(r.Package.Name)
	r.Package.RegistryIdentifier = toLowerCase(string(r.Package.Name[0]))
	if r.Suffix != "" {
		r.FullName = toCamelCase(r.Name + strings.ToTitle(r.Suffix))
		r.Suffix = "_" + toSnakeCase(r.Suffix)
	}
}
