package component

import (
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeMakefile(m filesystem.Manager) filesystem.File {
	template := "setup/scaffold/template/base/makefile"

	destination := m.RootDirectory

	resource := "script"

	content := "makefile"

	return base.New(base.ComponentInput{
		Identifier:           base.BuildNonModuleIdentifier(resource, content, destination),
		DestinationDirectory: destination,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			FileName:     "Makefile",
			Extension:    definition.NO_EXTENSION,
		})
}
