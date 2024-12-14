package bootstrap

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/organism/setup/constant"
	"github.com/charmingruby/bob/internal/component/organism/setup/data"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/scaffold"
)

func MakeEntry(m filesystem.Manager, module, repository string) filesystem.File {
	prepareDirectoriesForEntry(m)

	return base.New(base.ComponentInput{
		DestinationDirectory: m.EntryDirectory(),
	}).Componetize(
		scaffold.GENERATE_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.ENTRY_TEMPLATE,
			TemplateData: data.NewEntryData(m.RootPath(), module, repository),
			FileName:     "main",
		})
}

func prepareDirectoriesForEntry(m filesystem.Manager) {
	m.GenerateNestedDirectories(
		m.MainDirectory(),
		[]string{"cmd", m.ProjectName},
	)
}
