package fs

import (
	"os"
	"path/filepath"
)

func GenerateNestedDirectories(baseDir string, directories []string) error {
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
