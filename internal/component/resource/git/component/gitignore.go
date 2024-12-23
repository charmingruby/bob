package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeGitIgnore(m filesystem.Manager) filesystem.File {
	template := "resource/git/gitignore"

	return base.New(base.ComponentInput{
		DestinationDirectory: m.RootDirectory,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			FileName:     ".gitignore",
			Extension:    definition.NO_EXTENSION,
		})
}
