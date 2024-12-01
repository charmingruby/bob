package rest

import (
	"github.com/charmingruby/bob/internal/command/gen/atom"
	"github.com/charmingruby/bob/internal/command/gen/library"
	"github.com/charmingruby/bob/internal/command/gen/molecule/rest/rest_component"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/input"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunRest(m component.Manager) *cobra.Command {
	var (
		module string
	)

	cmd := &cobra.Command{
		Use:   "rest",
		Short: "Generates a new rest molecule",
		Run: func(cmd *cobra.Command, args []string) {
			if err := input.ValidateOnlyModuleCommandInput(module); err != nil {
				panic(err)
			}

			MakeRest(m, module)
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")

	return cmd
}

func MakeRest(m component.Manager, module string) {
	if err := filesystem.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{"transport", "rest"},
	); err != nil {
		panic(err)
	}

	if err := filesystem.GenerateFile(
		library.MakeValidatorComponent(m),
	); err != nil {
		panic(err)
	}

	if err := filesystem.GenerateMultipleDirectories(
		m.AppendToModuleDirectory(module, "transport/rest"),
		[]string{"endpoint", "dto"},
	); err != nil {
		panic(err)
	}

	if err := filesystem.GenerateMultipleDirectories(
		m.AppendToModuleDirectory(module, "transport/rest/dto"),
		[]string{"request", "response"},
	); err != nil {
		panic(err)
	}

	if err := filesystem.GenerateFile(
		rest_component.MakeRestRegistryComponent(
			m.AppendToModuleDirectory(module, "transport/rest/endpoint"),
			m.DependencyPath(module),
			module,
		)); err != nil {
		panic(err)
	}

	actioName := "ping"

	if err := filesystem.GenerateFile(atom.MakeHandlerComponent(
		m.SourceDirectory,
		module,
		actioName,
	)); err != nil {
		panic(err)
	}

	if err := filesystem.GenerateFile(rest_component.MakeRequest(
		m,
		module,
		actioName,
	)); err != nil {
		panic(err)
	}

	if err := filesystem.GenerateFile(rest_component.MakeResponse(
		m,
		module,
		actioName,
	)); err != nil {
		panic(err)
	}
}
