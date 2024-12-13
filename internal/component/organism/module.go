package organism

import (
	"github.com/charmingruby/bob/internal/component/atom"
	"github.com/charmingruby/bob/internal/component/molecule"
	"github.com/charmingruby/bob/internal/component/organism/module/component"
	"github.com/charmingruby/bob/internal/component/resource"
	"github.com/charmingruby/bob/internal/filesystem"
)

func MakeAndRunBaseModule(m filesystem.Manager, module string) {
	newModule := component.MakeBaseRegistry(m, module)
	if err := m.GenerateFile(newModule); err != nil {
		panic(err)
	}

	molecule.MakeAndRunCore(m, module)
	molecule.MakeAndRunRest(m, module)
}

func MakeAndRunModuleWithPostgresDatabase(m filesystem.Manager, module, modelName, tableName string) {
	newModule := component.MakeRegistryWithPostgresDatabase(m, module, modelName)
	if err := m.GenerateFile(newModule); err != nil {
		panic(err)
	}

	resource.MakeAndRunPostgresRepository(m, module, modelName, tableName, true)

	molecule.MakeAndRunCore(m, module)
	molecule.MakeAndRunRest(m, module)
}

func MakeAndRunModuleWithCustomDatabase(m filesystem.Manager, module, modelName, database string) {
	repo := atom.MakeUnimplementedRepository(m, module, modelName, database)
	if err := m.GenerateFile(repo); err != nil {
		panic(err)
	}

	molecule.MakeAndRunCore(m, module)
	molecule.MakeAndRunRest(m, module)
}
