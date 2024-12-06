package structure

import (
	"github.com/charmingruby/bob/internal/command/shared/component"
)

type RestRegistryData struct {
	Module     string
	SourcePath string
}

func NewRestRegistryData(module, sourcePath string) RestRegistryData {
	return RestRegistryData{
		Module:     component.ModuleFormat(module),
		SourcePath: sourcePath,
	}
}
