package filesystem

import (
	"fmt"

	"github.com/charmingruby/bob/config"
)

type Manager struct {
	Data            string
	SourceDirectory string
}

func New(config config.Configuration) Manager {
	var sourceDirectory = config.BaseConfiguration.RootDir + "/" + config.BaseConfiguration.SourceDir
	if config.BaseConfiguration.RootDir == "." {
		sourceDirectory = config.BaseConfiguration.SourceDir
	}

	return Manager{
		Data:            config.BaseConfiguration.BaseURL + "/" + config.BaseConfiguration.ProjectName,
		SourceDirectory: sourceDirectory,
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
