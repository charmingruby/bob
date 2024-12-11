package molecule

import (
	"github.com/charmingruby/bob/internal/component/library"
	"github.com/charmingruby/bob/internal/component/molecule/rest/component"
	"github.com/charmingruby/bob/internal/filesystem"
)

func MakeAndRunRest(m filesystem.Manager, module string) {
	if err := m.GenerateFile(
		library.MakeValidator(m),
	); err != nil {
		panic(err)
	}

	if err := m.GenerateFile(
		component.MakeRequestHelper(m),
	); err != nil {
		panic(err)
	}
	if err := m.GenerateFile(
		component.MakeBaseServerMiddleware(
			m,
		)); err != nil {
		panic(err)
	}

	if err := m.GenerateFile(
		component.MakeServer(
			m,
		)); err != nil {
		panic(err)
	}

	actioName := "ping"

	if err := m.GenerateFile(component.MakeHandler(
		m,
		module,
		actioName,
	)); err != nil {
		panic(err)
	}

	if err := m.GenerateFile(
		component.MakeHandlerRegistry(
			m.AppendToModuleDirectory(module, "transport/rest/endpoint"),
			m.DependencyPath(),
			module,
		)); err != nil {
		panic(err)
	}

	if err := m.GenerateFile(component.MakeRequestHelper(m)); err != nil {
		panic(err)
	}

	if err := m.GenerateFile(component.MakeResponseHelper(m)); err != nil {
		panic(err)
	}

	if err := m.GenerateFile(component.MakeRequest(
		m,
		module,
		actioName,
	)); err != nil {
		panic(err)
	}

	if err := m.GenerateFile(component.MakeResponse(
		m,
		module,
		actioName,
	)); err != nil {
		panic(err)
	}

}
