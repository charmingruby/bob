package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/organism/setup/constant"
	"github.com/charmingruby/bob/internal/component/organism/setup/data"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/scaffold"
)

func MakeEntry(m filesystem.Manager, module, repositoryModel string) filesystem.File {
	prepareDirectoriesForEntry(m)

	return base.New(base.ComponentInput{
		DestinationDirectory: m.EntryDirectory(),
	}).Componetize(
		scaffold.BOOTSTRAP_COMMAND,
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
