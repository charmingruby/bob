package container

import (
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/shared/resource/container/component"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func PerformDockerContainer(m filesystem.Manager, goVersion string) error {
	component := component.MakeContainer(m, goVersion)

	if err := m.GenerateFile(component); err != nil {
		return err
	}

	output.ComponentCreated(component.Identifier)

	return nil
}

func PerformDockerComposeWithPostgres(m filesystem.Manager) error {
	component := component.MakeComposeWithPostgres(m)

	if err := m.GenerateFile(component); err != nil {
		return err
	}

	output.ComponentCreated(component.Identifier)

	return nil
}

func PerformDockerCompose(m filesystem.Manager) error {
	component := component.MakeCompose(m)

	if err := m.GenerateFile(component); err != nil {
		return err
	}

	output.ComponentCreated(component.Identifier)

	return nil
}
