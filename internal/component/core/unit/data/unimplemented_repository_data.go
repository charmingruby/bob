package data

import (
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
)

type UnimplementedRepositoryData struct {
	SourcePath string
	Module     string
	Name       string
	Database   string
}

func NewUnimplementedRepositoryData(sourcePath, module, name, db string) UnimplementedRepositoryData {
	return UnimplementedRepositoryData{
		SourcePath: sourcePath,
		Module:     base.ModuleFormat(module),
		Name:       base.PublicNameFormat(name),
		Database:   base.ModuleFormat(db),
	}
}
