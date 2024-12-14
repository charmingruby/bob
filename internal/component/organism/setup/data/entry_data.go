package data

import "github.com/charmingruby/bob/internal/component/base"

type EntryData struct {
	RootPath                string
	Module                  string
	UpperCaseRepositoryName string
	LowerCaseRepositoryName string
}

func NewEntryData(
	rootPath, module, repositoryName string,
) EntryData {
	return EntryData{
		RootPath:                rootPath,
		Module:                  module,
		UpperCaseRepositoryName: base.PublicNameFormat(repositoryName),
		LowerCaseRepositoryName: base.PrivateNameFormat(repositoryName),
	}
}
