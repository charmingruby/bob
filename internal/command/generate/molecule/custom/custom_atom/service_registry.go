package custom_atom

import (
	"github.com/charmingruby/bob/internal/command/generate/atom"
	"github.com/charmingruby/bob/internal/command/generate/molecule/custom/custom_structure"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/fs"
)

func MakeServiceRegistryComponent(m component.Manager, module, name string) fs.File {
	return atom.MakeRegistryComponent(atom.RegistryParams{
		Module:               module,
		TemplateName:         "service_registry_with_repository",
		TemplateData:         custom_structure.NewServiceWithRepository(m.DependencyPath(module), module, name),
		RegistryName:         "service",
		DestinationDirectory: m.AppendToModuleDirectory(module, "/core/service"),
	})
}

func MakeIndependentServiceRegistryComponent(m component.Manager, module string) fs.File {
	return atom.MakeRegistryComponent(atom.RegistryParams{
		Module:               module,
		TemplateName:         "service_registry",
		RegistryName:         "service",
		DestinationDirectory: m.AppendToModuleDirectory(module, "/core/service"),
	})
}
