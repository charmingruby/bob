package data

import "github.com/charmingruby/bob/internal/component/base"

type ServiceWithRepositoryData struct {
	SourcePath            string
	Module                string
	RepositoryName        string
	PrivateRepositoryName string
}

func NewServiceWithRepositoryData(sourcePath, module, name string) ServiceWithRepositoryData {
	return ServiceWithRepositoryData{
		SourcePath:            sourcePath,
		Module:                base.ModuleFormat(module),
		RepositoryName:        base.PublicNameFormat(name),
		PrivateRepositoryName: base.PrivateNameFormat(name),
	}
}
