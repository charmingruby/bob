package rest

import (
	"github.com/charmingruby/bob/internal/command/gen/atom"
	restAtom "github.com/charmingruby/bob/internal/command/gen/molecule/rest/atom"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/input"
	"github.com/charmingruby/bob/internal/command/shared/fs"
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
	if err := fs.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{"transport", "rest"},
	); err != nil {
		panic(err)
	}

	if err := fs.GenerateMultipleDirectories(
		m.AppendToModuleDirectory(module, "transport/rest"),
		[]string{"endpoint", "dto"},
	); err != nil {
		panic(err)
	}

	if err := fs.GenerateMultipleDirectories(
		m.AppendToModuleDirectory(module, "transport/rest/dto"),
		[]string{"request", "response"},
	); err != nil {
		panic(err)
	}

	if err := fs.GenerateFile(
		restAtom.MakeRestRegistryComponent(
			m.AppendToModuleDirectory(module, "transport/rest/endpoint"),
			m.DependencyPath(module),
			module,
		)); err != nil {
		panic(err)
	}

	actioName := "ping"

	if err := fs.GenerateFile(atom.MakeHandlerComponent(
		m.SourceDirectory,
		module,
		actioName,
	)); err != nil {
		panic(err)
	}

	if err := fs.GenerateFile(restAtom.MakeRequest(
		m,
		module,
		actioName,
	)); err != nil {
		panic(err)
	}

	if err := fs.GenerateFile(restAtom.MakeResponse(
		m,
		module,
		actioName,
	)); err != nil {
		panic(err)
	}
}
