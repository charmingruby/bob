package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakePersistenceError(m filesystem.Manager) filesystem.File {
	template := "error/database/persistence"
	packageName := "database_err"

	return base.New(base.ComponentInput{
		Package: packageName,
		DestinationDirectory: definition.CustomErrPath(
			m.ModuleDirectory(definition.SHARED_MODULE),
			[]string{packageName},
		),
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			FileName:     "persistence",
			FileSuffix:   "err",
		})
}
