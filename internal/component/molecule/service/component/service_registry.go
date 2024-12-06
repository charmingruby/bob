package component

import (
	"github.com/charmingruby/bob/internal/component/atom"
	"github.com/charmingruby/bob/internal/component/molecule/constant"
	"github.com/charmingruby/bob/internal/component/molecule/service/data"
	"github.com/charmingruby/bob/internal/filesystem"
)

func MakeServiceRegistry(m filesystem.Manager, module, name string) filesystem.File {
	return atom.MakeRegistry(atom.RegistryParams{
		Package:              module,
		TemplateName:         constant.SERVICE_REGISTRY_WITH_REPOSITORY_TEMPLATE,
		TemplateData:         data.NewServiceWithRepositoryData(m.DependencyPath(), module, name),
		RegistryName:         "service",
		DestinationDirectory: m.AppendToModuleDirectory(module, "/core/service"),
	})
}

func MakeIndependentServiceRegistry(m filesystem.Manager, module string) filesystem.File {
	return atom.MakeRegistry(atom.RegistryParams{
		Package:              module,
		TemplateName:         constant.SERVICE_REGISTRY_TEMPLATE,
		RegistryName:         "service",
		DestinationDirectory: m.AppendToModuleDirectory(module, "/core/service"),
	})
}
