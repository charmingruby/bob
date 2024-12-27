package data

import "github.com/charmingruby/bob/internal/shared/definition/component/base"

type ModuleWithDatabaseData struct {
	SourcePath                string
	Module                    string
	Database                  string
	CapitalizedRepositoryName string
	LowerCaseRepositoryName   string
}

func NewModuleWithDatabaseData(sourcePath, module, database, repository string) ModuleWithDatabaseData {
	return ModuleWithDatabaseData{
		SourcePath:                sourcePath,
		Module:                    base.SnakeCaseFormat(module),
		Database:                  base.CamelCaseFormat(database),
		CapitalizedRepositoryName: base.CapitalizedFormat(repository),
		LowerCaseRepositoryName:   base.LowerCaseFormat(repository),
	}
}
