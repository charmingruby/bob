package library

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/scaffold"
)

func MakeValidator(m filesystem.Manager) filesystem.File {
	prepareDirectoriesForValidator(m, scaffold.SHARED_MODULE)

	return base.New(base.ComponentInput{
		Package:              scaffold.SHARED_MODULE,
		Name:                 "validator",
		DestinationDirectory: m.AppendToModuleDirectory(scaffold.SHARED_MODULE, "validation"),
	}).Componetize(base.ComponetizeInput{
		TemplateName: VALIDATION_TEMPLATE,
		FileName:     "validator",
	})
}

func prepareDirectoriesForValidator(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{module, "validation"},
	)
}
