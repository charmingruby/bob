package atom

import (
	"github.com/charmingruby/bob/internal/command/gen/atom"
	"github.com/charmingruby/bob/internal/command/gen/molecule/service/structure"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

func MakeServiceRegistryComponent(m component.Manager, module, name string) filesystem.File {
	return atom.MakeRegistryComponent(atom.RegistryParams{
		Module:               module,
		TemplateName:         "service_registry_with_repository",
		TemplateData:         structure.NewServiceWithRepository(m.DependencyPath(module), module, name),
		RegistryName:         "service",
		DestinationDirectory: m.AppendToModuleDirectory(module, "/core/service"),
	})
}

func MakeIndependentServiceRegistryComponent(m component.Manager, module string) filesystem.File {
	return atom.MakeRegistryComponent(atom.RegistryParams{
		Module:               module,
		TemplateName:         "service_registry",
		RegistryName:         "service",
		DestinationDirectory: m.AppendToModuleDirectory(module, "/core/service"),
	})
}
