package component

import (
	"github.com/charmingruby/bob/internal/component/context/rest"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

type entryData struct {
	RootPath                string
	Module                  string
	UpperCaseRepositoryName string
	LowerCaseRepositoryName string
}

func newEntryData(
	rootPath, module, repositoryName string,
) entryData {
	return entryData{
		RootPath:                rootPath,
		Module:                  module,
		UpperCaseRepositoryName: base.CapitalizedFormat(repositoryName),
		LowerCaseRepositoryName: base.LowerCaseFormat(repositoryName),
	}
}

func MakeEntry(m filesystem.Manager, module, repositoryModel string) filesystem.File {
	app := "api"

	prepareDirectoriesForEntry(m, app)

	template := rest.TemplatePath("template/base/entry")

	directory := m.ExecutableDirectory(app)

	resource := "exec"

	content := "application entry point"

	return base.New(base.ComponentInput{
		Identifier:           base.BuildNonModuleIdentifier(resource, content, m.ExecutableDirectory(app)),
		DestinationDirectory: directory,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			TemplateData: newEntryData(m.RootPath(), module, repositoryModel),
			FileName:     "main",
		})
}

func prepareDirectoriesForEntry(m filesystem.Manager, app string) {
	m.GenerateNestedDirectories(
		m.MainDirectory(),
		[]string{"cmd", app},
	)
}
