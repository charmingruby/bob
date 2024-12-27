package component

import (
	"github.com/charmingruby/bob/internal/component/shared/err"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeServiceError(m filesystem.Manager) filesystem.File {
	prepareDirectoriesForService(m)

	template := err.TemplatePath("core_err/service")

	destination := definition.CustomErrPath(
		m.ModuleDirectory(definition.SHARED_MODULE),
		[]string{definition.CORE_ERR_PACKAGE},
	)

	content := "service error"

	return base.New(base.ComponentInput{
		Identifier:           base.BuildIdentifier(definition.SHARED_MODULE, content, destination),
		Package:              definition.DATABASE_ERR_PACKAGE,
		DestinationDirectory: destination,
	}).Componetize(
		base.ComponetizeInput{
			TemplateName: template,
			FileName:     "service",
			FileSuffix:   "err",
		})
}

func prepareDirectoriesForService(m filesystem.Manager) {
	m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{definition.SHARED_MODULE, definition.CUSTOM_ERR_PACKAGE, definition.CORE_ERR_PACKAGE},
	)
}
