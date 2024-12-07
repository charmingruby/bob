package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/resource/database/postgres/constant"
	"github.com/charmingruby/bob/internal/component/resource/database/postgres/data"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/charmingruby/bob/internal/scaffold"
)

func MakePostgresRepository(m filesystem.Manager, module, model string) filesystem.File {
	return base.New(base.ComponentInput{
		Package:              constant.POSTGRES_PACKAGE,
		DestinationDirectory: scaffold.PersistencePath(m.ModuleDirectory(module), []string{constant.POSTGRES_REPOSITORY_PACKAGE}),
	}).Componetize(base.ComponetizeInput{
		TemplateName: constant.POSTGRES_REPOSITORY_TEMPLATE,
		TemplateData: data.NewPostgresRepositoryData(m.DependencyPath(), module, model),
		FileName:     model,
		FileSuffix:   "repository",
	})
}
