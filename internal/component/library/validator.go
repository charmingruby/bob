package library

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeValidator(m filesystem.Manager) filesystem.File {
	prepareDirectoriesForValidator(m, definition.SHARED_MODULE)

	template := "library/validator"

	return base.New(base.ComponentInput{
		Package:              definition.SHARED_MODULE,
		Name:                 "validator",
		DestinationDirectory: m.AppendToModuleDirectory(definition.SHARED_MODULE, "validation"),
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			FileName:     "validator",
		})
}

func prepareDirectoriesForValidator(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{module, "validation"},
	)
}
