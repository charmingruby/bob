package structure

import "github.com/charmingruby/bob/internal/command/shared/component/structure"

type RestRegistryData struct {
	Module          string
	SourceDirectory string
}

func NewRestRegistryData(module, sourceDirectory string) RestRegistryData {
	return RestRegistryData{
		Module:          structure.ModuleFormat(module),
		SourceDirectory: sourceDirectory,
	}
}
