package component

import (
	"github.com/charmingruby/bob/internal/component/core/unit"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

type serviceRegistryWithRepositoryData struct {
	SourcePath            string
	Module                string
	RepositoryName        string
	PrivateRepositoryName string
}

func newServiceRegistryWithRepositoryData(sourcePath, module, name string) serviceRegistryWithRepositoryData {
	return serviceRegistryWithRepositoryData{
		SourcePath:            sourcePath,
		Module:                base.ModuleFormat(module),
		RepositoryName:        base.PublicNameFormat(name),
		PrivateRepositoryName: base.PrivateNameFormat(name),
	}
}

func MakeServiceRegistry(m filesystem.Manager, module, name string) filesystem.File {
	template := "core/bundle/service/registry_with_repository"

	destination := m.AppendToModuleDirectory(module, "core/service")

	content := "service entrypoint"

	return unit.MakeRegistry(unit.RegistryParams{
		Identifier:           base.BuildIdentifier(module, content, destination),
		Package:              module,
		TemplateName:         template,
		TemplateData:         newServiceRegistryWithRepositoryData(m.DependencyPath(), module, name),
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
