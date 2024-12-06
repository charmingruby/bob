package structure

import (
	"github.com/charmingruby/bob/internal/command/shared/component"
)

type DependentPackageData struct {
	SourcePath string
	Module     string
	Name       string
}

func NewDependentPackageData(sourcePath, module, name string) DependentPackageData {
	return DependentPackageData{
		SourcePath: sourcePath,
		Module:     component.ModuleFormat(module),
		Name:       component.PublicNameFormat(name),
	}
}
