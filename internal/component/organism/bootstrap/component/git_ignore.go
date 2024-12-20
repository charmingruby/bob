package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/organism/bootstrap/constant"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeGitIgnore(m filesystem.Manager) filesystem.File {
	return base.New(base.ComponentInput{
		DestinationDirectory: m.RootDirectory,
	}).Componetize(
		definition.CREATE_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.GIT_IGNORE_TEMPLATE,
			FileName:     ".gitignore",
			Extension:    definition.NO_EXTENSION,
		})
}
