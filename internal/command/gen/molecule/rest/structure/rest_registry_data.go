package structure

import "github.com/charmingruby/bob/internal/command/shared/component/structure"

type RestRegistryData struct {
	Module     string
	SourcePath string
}

func NewRestRegistryData(module, sourcePath string) RestRegistryData {
	return RestRegistryData{
		Module:     structure.ModuleFormat(module),
		SourcePath: sourcePath,
	}
}
