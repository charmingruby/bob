package component

import (
	"fmt"

	"github.com/charmingruby/bob/config"
)

type Manager struct {
	Data            string
	SourceDirectory string
}

func NewManager(config config.Configuration) Manager {
	return Manager{
		Data:            config.BaseConfiguration.BaseURL + "/" + config.BaseConfiguration.ProjectName,
		SourceDirectory: config.BaseConfiguration.RootDir + "/" + config.BaseConfiguration.SourceDir,
	}
}

func (m *Manager) ModuleDirectory(module string) string {
	return m.SourceDirectory + "/" + module
}

func (m *Manager) AppendToModuleDirectory(module, path string) string {
	return m.ModuleDirectory(module) + "/" + path
}

func (m *Manager) DependencyPath(module string) string {
	return fmt.Sprintf("%s/%s", m.Data, m.SourceDirectory)
}

func ModulePath(sourceDirectory, module, path string) string {
	return sourceDirectory + "/" + module + "/" + path
}
