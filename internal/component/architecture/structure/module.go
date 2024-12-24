package structure

import (
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/architecture/bundle"
	"github.com/charmingruby/bob/internal/component/architecture/structure/module/component"
	"github.com/charmingruby/bob/internal/component/architecture/unit"
	"github.com/charmingruby/bob/internal/component/resource"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func PerformBaseModule(m filesystem.Manager, module, modelName string) error {
	newModule := component.MakeBaseRegistry(m, module)
	if err := m.GenerateFile(newModule); err != nil {
		return err
	}

	output.ComponentCreated(newModule.Identifier)

	if err := bundle.PerformCore(m, module, modelName); err != nil {
		return err
	}

	if err := bundle.PerformRest(m, module); err != nil {
		return err
	}

	return nil
}

func PerformModuleWithPostgresDatabase(m filesystem.Manager, module, modelName, tableName string) error {
	newModule := component.MakeRegistryWithPostgresDatabase(m, module, modelName)
	if err := m.GenerateFile(newModule); err != nil {
		return err
	}

	output.ComponentCreated(newModule.Identifier)

	if err := resource.PerformPostgresRepository(m, module, modelName, tableName, true); err != nil {
		return err
	}

	if err := bundle.PerformCore(m, module, modelName); err != nil {
		return err
	}

	if err := bundle.PerformRest(m, module); err != nil {
		return err
	}

	return nil
}

func PerformModuleWithCustomDatabase(m filesystem.Manager, module, modelName, database string) error {
	newModule := component.MakeRegistryWithCustomDatabase(m, module, modelName, database)
	if err := m.GenerateFile(newModule); err != nil {
		panic(err)
	}

	output.ComponentCreated(newModule.Identifier)

	repo := unit.MakeUnimplementedRepository(m, module, modelName, database)
	if err := m.GenerateFile(repo); err != nil {
		return err
	}

	output.ComponentCreated(repo.Identifier)

	if err := bundle.PerformCore(m, module, modelName); err != nil {
		return err
	}

	if err := bundle.PerformRest(m, module); err != nil {
		return err
	}

	return nil
}
