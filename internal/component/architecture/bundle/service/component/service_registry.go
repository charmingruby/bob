package component

import (
	"github.com/charmingruby/bob/internal/component/architecture/bundle/service/data"
	"github.com/charmingruby/bob/internal/component/architecture/unit"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeServiceRegistry(m filesystem.Manager, module, name string) filesystem.File {
	template := "architecture/bundle/service/registry_with_repository"

	return unit.MakeRegistry(unit.RegistryParams{
		Package:              module,
		TemplateName:         template,
		TemplateData:         data.NewServiceWithRepositoryData(m.DependencyPath(), module, name),
		RegistryName:         "service",
		DestinationDirectory: m.AppendToModuleDirectory(module, "/core/service"),
	})
}

func MakeIndependentServiceRegistry(m filesystem.Manager, module string) filesystem.File {
	template := "architecture/bundle/service/registry"

	return unit.MakeRegistry(unit.RegistryParams{
		Package:              module,
		TemplateName:         template,
		RegistryName:         "service",
		DestinationDirectory: m.AppendToModuleDirectory(module, "/core/service"),
	})
}
