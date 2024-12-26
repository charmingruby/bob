package handler

import (
	"github.com/charmingruby/bob/internal/cli/output"
	"github.com/charmingruby/bob/internal/component/context/rest/component"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func Perform(m filesystem.Manager, module, actionName string) error {
	components := []filesystem.File{
		component.MakeHandler(
			m,
			module,
			actionName,
		),
		component.MakeRequest(
			m,
			module,
			actionName,
		),
		component.MakeResponse(
			m,
			module,
			actionName,
		),
	}

	for _, f := range components {
		if err := m.GenerateFile(f); err != nil {
			return err
		}

		output.ComponentCreated(f.Identifier)
	}

	return nil
}
