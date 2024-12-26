package configure

import (
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/setup/configure/component"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func PerformConfigure() error {
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
