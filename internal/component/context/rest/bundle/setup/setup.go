package setup

import (
	"github.com/charmingruby/bob/internal/component/context/rest/component"
	"github.com/charmingruby/bob/internal/component/shared/library"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func Perform(m filesystem.Manager, module string) ([]filesystem.File, error) {
	actionName := "ping"

	components := []filesystem.File{
		library.MakeValidator(m),
		component.MakeBaseServerMiddleware(m),
		component.MakeServer(m),
		component.MakeHandler(
			m,
			module,
			actionName,
		),
		component.MakeHandlerRegistry(
			m.AppendToModuleDirectory(module, "transport/rest/endpoint"),
			m.DependencyPath(),
			module,
		),
		component.MakeRequestHelper(m),
		component.MakeResponseHelper(m),
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
			return nil, err
		}
	}

	return components, nil
}
