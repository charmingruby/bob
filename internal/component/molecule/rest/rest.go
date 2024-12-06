package rest

import (
	"github.com/charmingruby/bob/internal/component/library"
	"github.com/charmingruby/bob/internal/component/molecule/rest/component"
	"github.com/charmingruby/bob/internal/filesystem"
)

func MakeRest(m filesystem.Manager, module string) {
	if err := m.GenerateFile(
		library.MakeValidatorComponent(m),
	); err != nil {
		panic(err)
	}

	if err := m.GenerateFile(
		component.MakeRequestHelperComponent(m),
	); err != nil {
		panic(err)
	}
	if err := m.GenerateFile(
		component.MakeBaseServerMiddlewareComponent(
			m,
		)); err != nil {
		panic(err)
	}

	if err := m.GenerateFile(
		component.MakeServerComponent(
			m,
		)); err != nil {
		panic(err)
	}

	actioName := "ping"

	if err := m.GenerateFile(component.MakeHandlerComponent(
		m,
		module,
		actioName,
	)); err != nil {
		panic(err)
	}

	if err := m.GenerateFile(
		component.MakeHandlerRegistryComponent(
			m.AppendToModuleDirectory(module, "transport/rest/endpoint"),
			m.DependencyPath(),
			module,
		)); err != nil {
		panic(err)
	}

	if err := m.GenerateFile(component.MakeRequestHelperComponent(m)); err != nil {
		panic(err)
	}

	if err := m.GenerateFile(component.MakeResponseHelperComponent(m)); err != nil {
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
