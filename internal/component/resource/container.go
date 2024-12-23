package resource

import (
	"github.com/charmingruby/bob/internal/component/resource/container/component"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func PerformDockerContainer(m filesystem.Manager, goVersion string) {
	if err := m.GenerateFile(component.MakeContainer(m, goVersion)); err != nil {
		panic(err)
	}
}

func PerformDockerCompose(m filesystem.Manager) {
	if err := m.GenerateFile(component.MakeCompose(m)); err != nil {
		panic(err)
	}
}
