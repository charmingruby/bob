package library

import (
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/constant"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

func MakeValidatorComponent(m component.Manager) filesystem.File {
	if err := filesystem.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{constant.COMMON_MODULE, "validation"},
	); err != nil {
		panic(err)
	}

	return component.New(component.ComponentInput{
		Module:               constant.COMMON_MODULE,
		Name:                 "validator",
		DestinationDirectory: m.AppendToModuleDirectory(constant.COMMON_MODULE, "validation"),
	}).Componetize(component.ComponetizeInput{
		TemplateName: constant.VALIDATION_LIBRARY_TEMPLATE,
		FileName:     "validator",
	})
}
