package component

import (
	"github.com/charmingruby/bob/internal/component/shared/err"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakePersistenceError(m filesystem.Manager) filesystem.File {
	template := err.TemplatePath("database/persistence")

	destination := definition.CustomErrPath(
		m.ModuleDirectory(definition.SHARED_MODULE),
		[]string{definition.DATABASE_ERR_PACKAGE},
	)

	content := "persistence error"

	return base.New(base.ComponentInput{
		Identifier:           base.BuildIdentifier(definition.SHARED_MODULE, content, destination),
		Package:              definition.DATABASE_ERR_PACKAGE,
		DestinationDirectory: destination,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			FileName:     "persistence",
			FileSuffix:   "err",
		})
}
