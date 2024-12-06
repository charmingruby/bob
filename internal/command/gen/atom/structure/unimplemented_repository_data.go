package structure

import (
	"github.com/charmingruby/bob/internal/command/shared/component"
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
		Module:     component.ModuleFormat(module),
		Name:       component.PublicNameFormat(name),
		Database:   component.ModuleFormat(db),
	}
}
