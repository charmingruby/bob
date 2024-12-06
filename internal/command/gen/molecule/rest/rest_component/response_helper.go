package rest_component

import (
	"github.com/charmingruby/bob/internal/command/gen/molecule/rest/constant"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
	"github.com/charmingruby/bob/internal/command/shared/scaffold"
)

func MakeResponseHelperComponent(m filesystem.Manager) filesystem.File {
	if err := m.GenerateNestedDirectories(
		m.ModuleDirectory(scaffold.SHARED_MODULE),
		[]string{"transport", scaffold.REST_PACKAGE},
	); err != nil {
		panic(err)
	}

	return component.New(component.ComponentInput{
		Package: scaffold.REST_PACKAGE,
		DestinationDirectory: scaffold.TransportPath(
			m.ModuleDirectory(scaffold.SHARED_MODULE),
			scaffold.REST_PACKAGE,
			nil,
		),
	}).Componetize(component.ComponetizeInput{
		TemplateName: constant.REST_RESPONSE_HELPER_TEMPLATE,
		FileName:     "response",
	})
}
