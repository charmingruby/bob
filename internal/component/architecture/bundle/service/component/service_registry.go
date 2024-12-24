package component

import (
	"github.com/charmingruby/bob/internal/component/architecture/bundle/service/data"
	"github.com/charmingruby/bob/internal/component/architecture/unit"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeServiceRegistry(m filesystem.Manager, module, name string) filesystem.File {
	template := "architecture/bundle/service/registry_with_repository"

	destination := m.AppendToModuleDirectory(module, "core/service")

	content := "service entrypoint"

	return unit.MakeRegistry(unit.RegistryParams{
		Identifier:           base.BuildIdentifier(module, content, destination),
		Package:              module,
		TemplateName:         template,
		TemplateData:         data.NewServiceWithRepositoryData(m.DependencyPath(), module, name),
		RegistryName:         "service",
		DestinationDirectory: destination,
	})
}

func MakeIndependentServiceRegistry(m filesystem.Manager, module string) filesystem.File {
	template := "architecture/bundle/service/registry"

	destination := m.AppendToModuleDirectory(module, "core/service")

	content := "service entrypoint"

	return unit.MakeRegistry(unit.RegistryParams{
		Identifier:           base.BuildIdentifier(module, content, destination),
		Package:              module,
		TemplateName:         template,
		RegistryName:         "service",
		DestinationDirectory: destination,
	})
}
