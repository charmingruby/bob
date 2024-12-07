package filesystem

import (
	"fmt"

	"github.com/charmingruby/bob/config"
)

type Manager struct {
	Data             string
	RootDirectory    string
	SourceDirectory  string
	LibraryDirectory string
}

func New(config config.Configuration) Manager {
	var sourceDirectory = config.BaseConfiguration.RootDir + "/" + config.BaseConfiguration.SourceDir
	var libraryDirectory = config.BaseConfiguration.RootDir + "/" + config.BaseConfiguration.LibraryDir

	if config.BaseConfiguration.RootDir == "." {
		sourceDirectory = config.BaseConfiguration.SourceDir
		libraryDirectory = config.BaseConfiguration.LibraryDir
	}

	return Manager{
		Data:             config.BaseConfiguration.BaseURL + "/" + config.BaseConfiguration.ProjectName,
		RootDirectory:    config.BaseConfiguration.RootDir,
		SourceDirectory:  sourceDirectory,
		LibraryDirectory: libraryDirectory,
	}
}

func (m *Manager) MainDirectory() string {
	return m.RootDirectory
}

func (m *Manager) ModuleDirectory(module string) string {
	return m.SourceDirectory + "/" + module
}

func (m *Manager) AppendToModuleDirectory(module, path string) string {
	return m.ModuleDirectory(module) + "/" + path
}

func (m *Manager) ExternalLibraryDirectory(lib string) string {
	return m.LibraryDirectory + "/" + lib
}

func (m *Manager) DependencyPath() string {
	return fmt.Sprintf("%s/%s", m.Data, m.SourceDirectory)
}
