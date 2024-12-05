package rest_component

import (
	restConstant "github.com/charmingruby/bob/internal/command/gen/molecule/rest/constant"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/constant"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

const serverPath = "transport/rest"

func MakeServerComponent(m filesystem.Manager) filesystem.File {
	return component.New(component.ComponentInput{
		Module:               "rest",
		DestinationDirectory: m.AppendToModuleDirectory(constant.SHARED_MODULE, serverPath),
	}).Componetize(component.ComponetizeInput{
		TemplateName: restConstant.REST_SERVER,
		FileName:     "server",
	})
}
