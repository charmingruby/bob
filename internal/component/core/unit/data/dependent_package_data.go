package data

import (
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
)

type DependentPackageData struct {
	SourcePath string
	Module     string
	Name       string
}

func NewDependentPackageData(sourcePath, module, name string) DependentPackageData {
	return DependentPackageData{
		SourcePath: sourcePath,
		Module:     base.ModuleFormat(module),
		Name:       base.PublicNameFormat(name),
	}
}
