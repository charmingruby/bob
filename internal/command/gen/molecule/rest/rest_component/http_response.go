package rest_component

import (
	"github.com/charmingruby/bob/internal/command/gen/atom"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/constant"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

func MakeRestUtilComponent(m component.Manager) filesystem.File {
	pkg := "rest"

	if err := filesystem.GenerateNestedDirectories(
		m.ModuleDirectory(constant.COMMON_MODULE),
		[]string{"transport", pkg},
	); err != nil {
		panic(err)
	}

	component := atom.New(atom.ComponentInput{
		Module: pkg,
		DestinationDirectory: m.AppendToModuleDirectory(
			constant.COMMON_MODULE,
			"transport/rest"),
	})

	return atom.MakeCustomComponent(atom.CustomComponentInput{
		BaseComponent: *component,
		TemplateName:  constant.REST_HTTP_RESPONSE_TEMPLATE,
		TemplateData:  nil,
		FileName:      "response",
		FileSuffix:    "",
	})
}
