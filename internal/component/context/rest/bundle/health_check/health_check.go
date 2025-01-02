package health_check

import (
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func Perform(m filesystem.Manager) ([]filesystem.File, error) {
	prepareDirectories(m)

	components := []filesystem.File{
		MakeRegistry(m),
		MakeHandler(m),
	}

	for _, f := range components {
		if err := m.GenerateFile(f); err != nil {
			return nil, err
		}
	}

	return components, nil
}

func prepareDirectories(m filesystem.Manager) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(definition.SHARED_MODULE),
		[]string{definition.TRANSPORT_PACKAGE, definition.REST_PACKAGE, definition.HANDLER_PACKAGE},
	)
}
