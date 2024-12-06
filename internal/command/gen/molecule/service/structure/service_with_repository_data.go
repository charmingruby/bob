package structure

import "github.com/charmingruby/bob/internal/command/shared/component"

type ServiceWithRepositoryData struct {
	SourcePath            string
	Module                string
	RepositoryName        string
	PrivateRepositoryName string
}

func NewServiceWithRepositoryData(sourcePath, module, name string) ServiceWithRepositoryData {
	return ServiceWithRepositoryData{
		SourcePath:            sourcePath,
		Module:                component.ModuleFormat(module),
		RepositoryName:        component.PublicNameFormat(name),
		PrivateRepositoryName: component.PrivateNameFormat(name),
	}
}
