package organism

import (
	"github.com/charmingruby/bob/internal/component/atom"
	"github.com/charmingruby/bob/internal/component/molecule"
	"github.com/charmingruby/bob/internal/component/organism/module/component"
	"github.com/charmingruby/bob/internal/component/resource"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func PerformBaseModule(m filesystem.Manager, module, modelName string) {
	newModule := component.MakeBaseRegistry(m, module)
	if err := m.GenerateFile(newModule); err != nil {
		panic(err)
	}

	molecule.PerformCore(m, module, modelName)
	molecule.PerformRest(m, module)
}

func PerformModuleWithPostgresDatabase(m filesystem.Manager, module, modelName, tableName string) {
	newModule := component.MakeRegistryWithPostgresDatabase(m, module, modelName)
	if err := m.GenerateFile(newModule); err != nil {
		panic(err)
	}

	resource.PerformPostgresRepository(m, module, modelName, tableName, true)

	molecule.PerformCore(m, module, modelName)
	molecule.PerformRest(m, module)
}

func PerformModuleWithCustomDatabase(m filesystem.Manager, module, modelName, database string) {
	newModule := component.MakeRegistryWithCustomDatabase(m, module, modelName, database)
	if err := m.GenerateFile(newModule); err != nil {
		panic(err)
	}

	repo := atom.MakeUnimplementedRepository(m, module, modelName, database)
	if err := m.GenerateFile(repo); err != nil {
		panic(err)
	}

	molecule.PerformCore(m, module, modelName)
	molecule.PerformRest(m, module)
}
