package data

import "github.com/charmingruby/bob/internal/shared/definition/component/base"

type BaseModuleData struct {
	SourcePath string
	Module     string
}

func NewBaseModuleData(sourcePath, module string) BaseModuleData {
	return BaseModuleData{
		SourcePath: sourcePath,
		Module:     base.ModuleFormat(module),
	}
}
