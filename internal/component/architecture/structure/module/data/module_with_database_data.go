package data

import "github.com/charmingruby/bob/internal/component/base"

type ModuleWithDatabaseData struct {
	SourcePath          string
	Module              string
	Database            string
	UpperCaseRepository string
	LowerCaseRepository string
}

func NewModuleWithDatabaseData(sourcePath, module, database, repository string) ModuleWithDatabaseData {
	return ModuleWithDatabaseData{
		SourcePath:          sourcePath,
		Module:              base.ModuleFormat(module),
		Database:            base.PackagePathFormat(database),
		UpperCaseRepository: base.PublicNameFormat(repository),
		LowerCaseRepository: base.PrivateNameFormat(repository),
	}
}
