package rest

import (
	"github.com/charmingruby/bob/internal/command/gen/library"
	"github.com/charmingruby/bob/internal/command/gen/molecule/rest/rest_component"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
	"github.com/charmingruby/bob/internal/command/shared/scaffold"
)

func MakeRest(m filesystem.Manager, module string) {
	if err := m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{"transport", "rest"},
	); err != nil {
		panic(err)
	}

	if err := m.GenerateFile(
		library.MakeValidatorComponent(m),
	); err != nil {
		panic(err)
	}

	if err := m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{scaffold.SHARED_MODULE, "transport", "rest"},
	); err != nil {
		panic(err)
	}

	if err := m.GenerateFile(
		rest_component.MakeRequestHelperComponent(m),
	); err != nil {
		panic(err)
	}

	if err := m.GenerateMultipleDirectories(
		m.AppendToModuleDirectory(module, "transport/rest"),
		[]string{"endpoint", "dto"},
	); err != nil {
		panic(err)
	}

	if err := m.GenerateMultipleDirectories(
		m.AppendToModuleDirectory(module, "transport/rest/dto"),
		[]string{"request", "response"},
	); err != nil {
		panic(err)
	}

	if err := m.GenerateFile(
		rest_component.MakeBaseServerMiddlewareComponent(
			m,
		)); err != nil {
		panic(err)
	}

	if err := m.GenerateFile(
		rest_component.MakeServerComponent(
			m,
		)); err != nil {
		panic(err)
	}

	if err := m.GenerateFile(
		rest_component.MakeHandlerRegistryComponent(
			m.AppendToModuleDirectory(module, "transport/rest/endpoint"),
			m.DependencyPath(),
			module,
		)); err != nil {
		panic(err)
	}

	actioName := "ping"

	if err := m.GenerateFile(rest_component.MakeRequestHelperComponent(m)); err != nil {
		panic(err)
	}

	if err := m.GenerateFile(rest_component.MakeResponseHelperComponent(m)); err != nil {
		panic(err)
	}

	if err := m.GenerateFile(rest_component.MakeRequest(
		m,
		module,
		actioName,
	)); err != nil {
		panic(err)
	}

	if err := m.GenerateFile(rest_component.MakeResponse(
		m,
		module,
		actioName,
	)); err != nil {
		panic(err)
	}

	if err := m.GenerateFile(rest_component.MakeHandlerComponent(
		m,
		module,
		actioName,
	)); err != nil {
		panic(err)
	}
}