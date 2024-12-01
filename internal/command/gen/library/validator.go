package library

import (
	"github.com/charmingruby/bob/internal/command/gen/atom"
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

	component := atom.New(atom.ComponentInput{
		Module:               constant.COMMON_MODULE,
		Name:                 "validator",
		Suffix:               "",
		DestinationDirectory: m.AppendToModuleDirectory(constant.COMMON_MODULE, "validation"),
		HasTest:              false,
	})

	return atom.MakeCustomComponent(
		atom.CustomComponentInput{
			BaseComponent: *component,
			TemplateName:  constant.VALIDATION_LIBRARY_TEMPLATE,
			TemplateData:  nil,
			FileName:      component.Name,
			FileSuffix:    "",
		},
	)
}
