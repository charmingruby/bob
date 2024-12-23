package component

import (
	"github.com/charmingruby/bob/internal/component/setup/scaffold/template/postgres/data"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeEntry(m filesystem.Manager, module, repositoryModel string) filesystem.File {
	app := "api"

	prepareDirectoriesForEntry(m, app)

	template := "setup/scaffold/template/base/entry"

	directory := m.EntryDirectory(app)

	resource := "exec"

	content := "application entry point"

	return base.New(base.ComponentInput{
		Identifier:           base.BuildNonModuleIdentifier(resource, content, m.EntryDirectory(app)),
		DestinationDirectory: directory,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			TemplateData: data.NewEntryData(m.RootPath(), module, repositoryModel),
			FileName:     "main",
		})
}

func prepareDirectoriesForEntry(m filesystem.Manager, app string) {
	m.GenerateNestedDirectories(
		m.MainDirectory(),
		[]string{"cmd", app},
	)
}
