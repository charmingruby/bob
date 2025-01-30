package component

import (
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeBobConfig(m filesystem.Manager) filesystem.File {
	template := "setup/schema/bob_config"

	destination := m.MainDirectory()

	content := "bob config file"

	return base.New(base.ComponentInput{
		Identifier:           base.BuildBobIdentifier(content, destination),
		DestinationDirectory: destination,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			FileName:     "bob",
			Extension:    definition.YML_EXTENSION,
		})
}
