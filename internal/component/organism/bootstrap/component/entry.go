package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/organism/bootstrap/constant"
	"github.com/charmingruby/bob/internal/component/organism/bootstrap/data"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/shared"
)

func MakeEntry(m filesystem.Manager, module, repositoryModel string) filesystem.File {
	prepareDirectoriesForEntry(m)

	return base.New(base.ComponentInput{
		DestinationDirectory: m.EntryDirectory("api"),
	}).Componetize(
		shared.BOOTSTRAP_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.ENTRY_TEMPLATE,
			TemplateData: data.NewEntryData(m.RootPath(), module, repositoryModel),
			FileName:     "main",
		})
}

func prepareDirectoriesForEntry(m filesystem.Manager) {
	m.GenerateNestedDirectories(
		m.MainDirectory(),
		[]string{"cmd", m.ProjectName},
	)
}
