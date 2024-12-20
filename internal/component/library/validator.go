package library

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/library/constant"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/shared"
)

func MakeValidator(m filesystem.Manager) filesystem.File {
	prepareDirectoriesForValidator(m, shared.SHARED_MODULE)

	return base.New(base.ComponentInput{
		Package:              shared.SHARED_MODULE,
		Name:                 "validator",
		DestinationDirectory: m.AppendToModuleDirectory(shared.SHARED_MODULE, "validation"),
	}).Componetize(
		shared.ADD_COMMAND,
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
