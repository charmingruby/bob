package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeMakefile(m filesystem.Manager) filesystem.File {
	template := "setup/scaffold/template/postgres/makefile"

	return base.New(base.ComponentInput{
		DestinationDirectory: m.RootDirectory,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			FileName:     "Makefile",
			Extension:    definition.NO_EXTENSION,
		})
}
