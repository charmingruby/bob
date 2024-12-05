package library

import (
	libConstant "github.com/charmingruby/bob/internal/command/gen/library/constant"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/constant"

	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

func MakeValidatorComponent(m filesystem.Manager) filesystem.File {
	if err := m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{constant.SHARED_MODULE, "validation"},
	); err != nil {
		panic(err)
	}

	return component.New(component.ComponentInput{
		Module:               constant.SHARED_MODULE,
		Name:                 "validator",
		DestinationDirectory: m.AppendToModuleDirectory(constant.SHARED_MODULE, "validation"),
	}).Componetize(component.ComponetizeInput{
		TemplateName: libConstant.VALIDATION_TEMPLATE,
		FileName:     "validator",
	})
}
