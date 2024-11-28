package custom_structure

import "github.com/charmingruby/bob/internal/command/shared/component/structure"

type ServiceWithRepositoryData struct {
	SourcePath            string
	Module                string
	RepositoryName        string
	PrivateRepositoryName string
}

func NewServiceWithRepository(sourcePath, module, name string) ServiceWithRepositoryData {
	return ServiceWithRepositoryData{
		SourcePath:            sourcePath,
		Module:                structure.ModuleFormat(module),
		RepositoryName:        structure.PublicNameFormat(name),
		PrivateRepositoryName: structure.PrivateNameFormat(name),
	}
}
