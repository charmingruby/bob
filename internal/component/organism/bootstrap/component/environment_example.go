package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/organism/bootstrap/constant"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/shared"
)

func MakeEnvironmentExample(m filesystem.Manager) filesystem.File {
	return base.New(base.ComponentInput{
		DestinationDirectory: m.RootDirectory,
	}).Componetize(
		shared.BOOTSTRAP_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.ENVIRONMENT_EXAMPLE_TEMPLATE,
			FileName:     ".env.example",
			Extension:    shared.NO_EXTENSION,
		})
}
