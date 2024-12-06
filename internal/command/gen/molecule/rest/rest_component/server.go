package rest_component

import (
	"github.com/charmingruby/bob/internal/command/gen/molecule/rest/constant"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
	"github.com/charmingruby/bob/internal/command/shared/scaffold"
)

func MakeServerComponent(m filesystem.Manager) filesystem.File {
	return component.New(component.ComponentInput{
		Package: scaffold.REST_PACKAGE,
		//DestinationDirectory: m.AppendToModuleDirectory(scaffold.SHARED_MODULE, path),
		DestinationDirectory: scaffold.TransportPath(
			m.ModuleDirectory(scaffold.SHARED_MODULE),
			scaffold.REST_PACKAGE,
			nil,
		),
	}).Componetize(component.ComponetizeInput{
		TemplateName: constant.REST_SERVER,
		FileName:     "server",
	})
}
