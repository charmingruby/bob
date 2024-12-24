package library

import (
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeValidator(m filesystem.Manager) filesystem.File {
	prepareDirectoriesForValidator(m, definition.SHARED_MODULE)

	template := "library/validator"

	destination := m.AppendToModuleDirectory(definition.SHARED_MODULE, "validation")

	content := "validation logic"

	return base.New(base.ComponentInput{
		Identifier:           base.BuildIdentifier(definition.SHARED_MODULE, content, destination),
		Package:              definition.SHARED_MODULE,
		Name:                 "validator",
		DestinationDirectory: destination,
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
