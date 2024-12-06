package structure

import "github.com/charmingruby/bob/internal/command/shared/component/structure"

type UnimplementedRepositoryData struct {
	SourcePath string
	Module     string
	Name       string
	Database   string
}

func NewUnimplementedRepositoryData(sourcePath, module, name, db string) UnimplementedRepositoryData {
	return UnimplementedRepositoryData{
		SourcePath: sourcePath,
		Module:     structure.ModuleFormat(module),
		Name:       structure.PublicNameFormat(name),
		Database:   structure.ModuleFormat(db),
	}
}
