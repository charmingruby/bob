package rest_component

import (
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

	return component.New(component.ComponentInput{
		Module: pkg,
		DestinationDirectory: m.AppendToModuleDirectory(
			constant.COMMON_MODULE,
			"transport/rest"),
	}).Componetize(component.ComponetizeInput{
		TemplateName: constant.REST_HTTP_RESPONSE_TEMPLATE,
		FileName:     "response",
	})
}
