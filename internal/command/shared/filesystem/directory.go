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

func (f *Manager) GenerateDirectory(baseDir, directory string) error {
	currentDir, err := getDestinationDirectory(baseDir)
	if err != nil {
		return err
	}

	newDir := filepath.Join(currentDir, directory)
	if _, err := os.Stat(newDir); os.IsNotExist(err) {
		if err := os.Mkdir(newDir, os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}

func (f *Manager) GenerateNestedDirectories(baseDir string, directories []string) error {
	currentDir, err := getDestinationDirectory(baseDir)
	if err != nil {
		return err
	}

	var currentPath = currentDir
	for _, directory := range directories {
		currentPath = filepath.Join(currentPath, directory)
		if _, err := os.Stat(currentPath); os.IsNotExist(err) {
			if err := os.Mkdir(currentPath, os.ModePerm); err != nil {
				return err
			}
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
		if _, err := os.Stat(newDir); os.IsNotExist(err) {
			if err := os.Mkdir(newDir, os.ModePerm); err != nil {
				return err
			}
		}
	}
	return nil
}
