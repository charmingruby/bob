package data

import "github.com/charmingruby/bob/internal/shared/definition/component/base"

type RestRegistryData struct {
	Module     string
	SourcePath string
}

func NewRestRegistryData(module, sourcePath string) RestRegistryData {
	return RestRegistryData{
		Module:     base.ModuleFormat(module),
		SourcePath: sourcePath,
	}
}
