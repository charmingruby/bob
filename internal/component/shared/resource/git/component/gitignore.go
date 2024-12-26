package component

import (
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeGitIgnore(m filesystem.Manager) filesystem.File {
	template := "resource/git/gitignore"

	destination := m.MainDirectory()

	resource := "git"

	content := "gitignore"

	return base.New(base.ComponentInput{
		Identifier:           base.BuildNonModuleIdentifier(resource, content, destination),
		DestinationDirectory: destination,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			FileName:     ".gitignore",
			Extension:    definition.NO_EXTENSION,
		})
}
