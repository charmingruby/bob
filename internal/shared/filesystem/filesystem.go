package filesystem

import (
	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/shared/definition"
)

type Manager struct {
	ProjectName      string
	Data             string
	RootDirectory    string
	SourceDirectory  string
	LibraryDirectory string
}

func New(config *config.Configuration) Manager {
	var sourceDirectory = config.BaseConfiguration.RootDir + "/" + definition.SOURCE_DIR
	var libraryDirectory = config.BaseConfiguration.RootDir + "/" + definition.LIB_DIR

	if config.BaseConfiguration.RootDir == "." {
		sourceDirectory = definition.SOURCE_DIR
		libraryDirectory = definition.LIB_DIR
	}

	return Manager{
		ProjectName:      config.BaseConfiguration.ProjectName,
		Data:             config.BaseConfiguration.BaseURL + "/" + config.BaseConfiguration.ProjectName,
		RootDirectory:    config.BaseConfiguration.RootDir,
		SourceDirectory:  sourceDirectory,
		LibraryDirectory: libraryDirectory,
	}
}

func (m *Manager) MainDirectory() string {
	return m.RootDirectory
}

func (m *Manager) ExecutableDirectory(project string) string {
	return m.MainDirectory() + "/cmd/" + project
}

func (m *Manager) ModuleDirectory(module string) string {
	return m.SourceDirectory + "/" + module
}

func (m *Manager) ExternalLibraryDirectory(lib string) string {
	return m.LibraryDirectory + "/" + lib
}

func (m *Manager) AppendToModuleDirectory(module, path string) string {
	return m.ModuleDirectory(module) + "/" + path
}
