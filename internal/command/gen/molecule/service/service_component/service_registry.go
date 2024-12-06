package service_component

import (
	"github.com/charmingruby/bob/internal/command/gen/atom"
	"github.com/charmingruby/bob/internal/command/gen/molecule/service/constant"
	"github.com/charmingruby/bob/internal/command/gen/molecule/service/structure"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

func MakeServiceRegistryComponent(m filesystem.Manager, module, name string) filesystem.File {
	return atom.MakeRegistryComponent(atom.RegistryParams{
		Package:              module,
		TemplateName:         constant.SERVICE_REGISTRY_WITH_REPOSITORY_TEMPLATE,
		TemplateData:         structure.NewServiceWithRepositoryData(m.DependencyPath(), module, name),
		RegistryName:         "service",
		DestinationDirectory: m.AppendToModuleDirectory(module, "/core/service"),
	})
}

func MakeIndependentServiceRegistryComponent(m filesystem.Manager, module string) filesystem.File {
	return atom.MakeRegistryComponent(atom.RegistryParams{
		Package:              module,
		TemplateName:         constant.SERVICE_REGISTRY_TEMPLATE,
		RegistryName:         "service",
		DestinationDirectory: m.AppendToModuleDirectory(module, "/core/service"),
	})
}
