package schema

import (
	"github.com/charmingruby/bob/internal/component/setup/schema/component"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func Perform() (filesystem.File, error) {
	manager := filesystem.Manager{
		ProjectName:      "bob-project",
		Data:             "github.com/your-user",
		RootDirectory:    ".",
		SourceDirectory:  "internal",
		LibraryDirectory: "pkg",
	}

	component := component.MakeBobConfig(manager)

	if err := manager.GenerateFile(component); err != nil {
		return filesystem.File{}, err
	}

	return component, nil
}
