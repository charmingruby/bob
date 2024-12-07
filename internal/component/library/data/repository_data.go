package data

import "github.com/charmingruby/bob/internal/component/base"

type PostgresRepositoryData struct {
	Module         string
	SourcePath     string
	LowerCaseModel string
	UpperCaseModel string
}

func NewPostgresRepositoryData(sourcePath, module, model string) PostgresRepositoryData {
	return PostgresRepositoryData{
		Module:         base.ModuleFormat(module),
		SourcePath:     sourcePath,
		LowerCaseModel: base.PrivateNameFormat(model),
		UpperCaseModel: base.PublicNameFormat(model),
	}
}
