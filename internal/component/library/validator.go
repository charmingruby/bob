package library

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/library/constant"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeValidator(m filesystem.Manager) filesystem.File {
	prepareDirectoriesForValidator(m, definition.SHARED_MODULE)

	return base.New(base.ComponentInput{
		Package:              definition.SHARED_MODULE,
		Name:                 "validator",
		DestinationDirectory: m.AppendToModuleDirectory(definition.SHARED_MODULE, "validation"),
	}).Componetize(
		definition.ADD_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.VALIDATION_TEMPLATE,
			FileName:     "validator",
		})
}

func prepareDirectoriesForValidator(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{module, "validation"},
	)
}
