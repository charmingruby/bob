package component

import (
	"github.com/charmingruby/bob/internal/component/atom"
	"github.com/charmingruby/bob/internal/component/molecule"
	"github.com/charmingruby/bob/internal/component/molecule/service/data"
	"github.com/charmingruby/bob/internal/filesystem"
)

func MakeServiceRegistryComponent(m filesystem.Manager, module, name string) filesystem.File {
	return atom.MakeRegistryComponent(atom.RegistryParams{
		Package:              module,
		TemplateName:         molecule.SERVICE_REGISTRY_WITH_REPOSITORY_TEMPLATE,
		TemplateData:         data.NewServiceWithRepositoryData(m.DependencyPath(), module, name),
		RegistryName:         "service",
		DestinationDirectory: m.AppendToModuleDirectory(module, "/core/service"),
	})
}

func MakeIndependentServiceRegistryComponent(m filesystem.Manager, module string) filesystem.File {
	return atom.MakeRegistryComponent(atom.RegistryParams{
		Package:              module,
		TemplateName:         molecule.SERVICE_REGISTRY_TEMPLATE,
		RegistryName:         "service",
		DestinationDirectory: m.AppendToModuleDirectory(module, "/core/service"),
	})
}
