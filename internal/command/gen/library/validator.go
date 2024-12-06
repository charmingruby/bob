package library

import (
	"github.com/charmingruby/bob/internal/command/gen/library/constant"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/scaffold"

	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

func MakeValidatorComponent(m filesystem.Manager) filesystem.File {
	if err := m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{scaffold.SHARED_MODULE, "validation"},
	); err != nil {
		panic(err)
	}

	return component.New(component.ComponentInput{
		Package:              scaffold.SHARED_MODULE,
		Name:                 "validator",
		DestinationDirectory: m.AppendToModuleDirectory(scaffold.SHARED_MODULE, "validation"),
	}).Componetize(component.ComponetizeInput{
		TemplateName: constant.VALIDATION_TEMPLATE,
		FileName:     "validator",
	})
}
