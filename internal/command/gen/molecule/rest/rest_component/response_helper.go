package rest_component

import (
	restConstant "github.com/charmingruby/bob/internal/command/gen/molecule/rest/constant"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/constant"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

func MakeResponseHelperComponent(m filesystem.Manager) filesystem.File {
	pkg := "rest"

	if err := m.GenerateNestedDirectories(
		m.ModuleDirectory(constant.SHARED_MODULE),
		[]string{"transport", pkg},
	); err != nil {
		panic(err)
	}

	return component.New(component.ComponentInput{
		Module: pkg,
		DestinationDirectory: m.AppendToModuleDirectory(
			constant.SHARED_MODULE,
			"transport/rest",
		),
	}).Componetize(component.ComponetizeInput{
		TemplateName: restConstant.REST_RESPONSE_HELPER_TEMPLATE,
		FileName:     "response",
	})
}
