package atom

import (
	"github.com/charmingruby/bob/internal/command/gen/atom/constant"
	"github.com/charmingruby/bob/internal/command/gen/atom/structure"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

var repositoryPath = "core/repository"

func MakeRepositoryComponent(m filesystem.Manager, module, name string) filesystem.File {
	return component.New(component.ComponentInput{
		DestinationDirectory: filesystem.ModulePath(m.SourceDirectory, module, repositoryPath),
		Module:               module,
		Name:                 name,
		Suffix:               "repository",
	}).Componetize(component.ComponetizeInput{
		TemplateName: constant.REPOSITORY_TEMPLATE,
		TemplateData: structure.NewDependentPackageData(m.DependencyPath(), module, name),
		FileName:     name,
		FileSuffix:   "repository",
	})
}
