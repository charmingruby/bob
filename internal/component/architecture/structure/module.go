package structure

import (
	"github.com/charmingruby/bob/internal/component/architecture/bundle"
	"github.com/charmingruby/bob/internal/component/architecture/structure/module/component"
	"github.com/charmingruby/bob/internal/component/architecture/unit"
	"github.com/charmingruby/bob/internal/component/resource"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func PerformBaseModule(m filesystem.Manager, module, modelName string) {
	newModule := component.MakeBaseRegistry(m, module)
	if err := m.GenerateFile(newModule); err != nil {
		panic(err)
	}

	bundle.PerformCore(m, module, modelName)
	bundle.PerformRest(m, module)
}

func PerformModuleWithPostgresDatabase(m filesystem.Manager, module, modelName, tableName string) {
	newModule := component.MakeRegistryWithPostgresDatabase(m, module, modelName)
	if err := m.GenerateFile(newModule); err != nil {
		panic(err)
	}

	resource.PerformPostgresRepository(m, module, modelName, tableName, true)

	bundle.PerformCore(m, module, modelName)
	bundle.PerformRest(m, module)
}

func PerformModuleWithCustomDatabase(m filesystem.Manager, module, modelName, database string) {
	newModule := component.MakeRegistryWithCustomDatabase(m, module, modelName, database)
	if err := m.GenerateFile(newModule); err != nil {
		panic(err)
	}

	repo := unit.MakeUnimplementedRepository(m, module, modelName, database)
	if err := m.GenerateFile(repo); err != nil {
		panic(err)
	}

	bundle.PerformCore(m, module, modelName)
	bundle.PerformRest(m, module)
}
