package setup

import (
	"github.com/charmingruby/bob/internal/component/setup/configure/component"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func PerformConfigure() {
	manager := filesystem.Manager{
		ProjectName:      "bob-project",
		Data:             "github.com/your-user",
		RootDirectory:    ".",
		SourceDirectory:  "internal",
		LibraryDirectory: "pkg",
	}

	component := component.MakeConfigure(manager)

	if err := manager.GenerateFile(component); err != nil {
		panic(err)
	}
}
