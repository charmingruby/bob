package schema

import (
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/config/schema/component"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func Perform() error {
	manager := filesystem.Manager{
		ProjectName:      "bob-project",
		Data:             "github.com/your-user",
		RootDirectory:    ".",
		SourceDirectory:  "internal",
		LibraryDirectory: "pkg",
	}

	component := component.MakeBobConfig(manager)

	if err := manager.GenerateFile(component); err != nil {
		return err
	}

	output.ComponentCreated(component.Identifier)

	return nil
}
