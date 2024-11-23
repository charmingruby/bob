package component

import "github.com/charmingruby/bob/internal/command/shared/formatter"

type Package struct {
	Name               string
	RegistryIdentifier string
	Registry           string
}

func (p *Package) Format() {
	p.Name = formatter.ToSnakeCase(p.Name)
	p.Registry = formatter.ToCamelCase(p.Name)
	p.RegistryIdentifier = formatter.ToLowerCase(string(p.Name[0]))
}
