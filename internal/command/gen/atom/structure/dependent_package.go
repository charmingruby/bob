package structure

import "github.com/charmingruby/bob/internal/command/shared/component/structure"

type DependentPackageData struct {
	SourcePath string
	Module     string
	Name       string
}

func NewDependentPackageData(sourcePath, module, name string) DependentPackageData {
	return DependentPackageData{
		SourcePath: sourcePath,
		Module:     structure.ModuleFormat(module),
		Name:       structure.PublicNameFormat(name),
	}
}
