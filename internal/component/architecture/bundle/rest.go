package bundle

import (
	"github.com/charmingruby/bob/internal/component/architecture/bundle/rest/component"
	"github.com/charmingruby/bob/internal/component/library"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func PerformRest(m filesystem.Manager, module string) {
	actioName := "ping"

	components := []filesystem.File{
		library.MakeValidator(m),
		component.MakeRequestHelper(m),
		component.MakeBaseServerMiddleware(m),
		component.MakeServer(m),
		component.MakeHandler(
			m,
			module,
			actioName,
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
			actioName,
		),
		component.MakeResponse(
			m,
			module,
			actioName,
		),
	}

	for _, f := range components {
		if err := m.GenerateFile(f); err != nil {
			panic(err)
		}
	}
}
