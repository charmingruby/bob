package filesystem

import (
	"os"
	"path/filepath"
)

func getDestinationDirectory(baseDir string) (string, error) {
	absPath, err := filepath.Abs(baseDir)
	if err != nil {
		return "", err
	}
	return absPath, nil
}

func ensureDirectoryExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func (f *Manager) GenerateDirectory(baseDir, directory string) error {
	currentDir, err := getDestinationDirectory(baseDir)
	if err != nil {
		return err
	}

	newDir := filepath.Join(currentDir, directory)
	return ensureDirectoryExists(newDir)
}

func (f *Manager) GenerateNestedDirectories(baseDir string, directories []string) error {
	currentDir, err := getDestinationDirectory(baseDir)
	if err != nil {
		return err
	}

	var currentPath = currentDir
	for _, directory := range directories {
		currentPath = filepath.Join(currentPath, directory)
		if err := ensureDirectoryExists(currentPath); err != nil {
			return err
		}
	}
	return nil
}

func (f *Manager) GenerateMultipleDirectories(baseDir string, directories []string) error {
	currentDir, err := getDestinationDirectory(baseDir)
	if err != nil {
		return err
	}

	for _, directory := range directories {
		newDir := filepath.Join(currentDir, directory)
		if err := ensureDirectoryExists(newDir); err != nil {
			return err
		}
	}
	return nil
}
