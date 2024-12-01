package filesystem

import (
	"fmt"

	"github.com/charmingruby/bob/config"
)

type FileSystem struct {
	Data            string
	SourceDirectory string
}

func New(config config.Configuration) FileSystem {
	var sourceDirectory = config.BaseConfiguration.RootDir + "/" + config.BaseConfiguration.SourceDir
	if config.BaseConfiguration.RootDir == "." {
		sourceDirectory = config.BaseConfiguration.SourceDir
	}

	return FileSystem{
		Data:            config.BaseConfiguration.BaseURL + "/" + config.BaseConfiguration.ProjectName,
		SourceDirectory: sourceDirectory,
	}
}

func (m *FileSystem) ModuleDirectory(module string) string {
	return m.SourceDirectory + "/" + module
}

func (m *FileSystem) AppendToModuleDirectory(module, path string) string {
	return m.ModuleDirectory(module) + "/" + path
}

func (m *FileSystem) DependencyPath(module string) string {
	return fmt.Sprintf("%s/%s", m.Data, m.SourceDirectory)
}

func ModulePath(sourceDirectory, module, path string) string {
	return sourceDirectory + "/" + module + "/" + path
}
