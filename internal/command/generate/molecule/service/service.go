package service

import (
	"fmt"

	"github.com/charmingruby/bob/internal/command/generate/atom"
	serviceAtom "github.com/charmingruby/bob/internal/command/generate/molecule/service/atom"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/input"
	"github.com/charmingruby/bob/internal/command/shared/fs"
	"github.com/spf13/cobra"
)

func RunService(m component.Manager) *cobra.Command {
	var (
		module string
		repo   string
	)

	cmd := &cobra.Command{
		Use:   "service",
		Short: "Generates a new service molecule",
		Run: func(cmd *cobra.Command, args []string) {
			if err := fs.GenerateNestedDirectories(
				fmt.Sprintf("%s/%s", m.SourceDirectory, module),
				[]string{"core", "service"},
			); err != nil {
				panic(err)
			}

			if err := input.ValidateOnlyModuleCommandInput(module); err != nil {
				panic(err)
			}

			MakeService(m, repo, module)
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&repo, "repo", "r", "", "repository dependency")

	return cmd
}

func MakeService(m component.Manager, repo string, module string) {
	hasRepo := repo != ""

	if !hasRepo {
		if err := fs.GenerateFile(serviceAtom.MakeIndependentServiceRegistryComponent(
			m,
			module,
		)); err != nil {
			panic(err)
		}
	} else {
		if err := fs.GenerateFile(serviceAtom.MakeServiceRegistryComponent(
			m,
			module,
			repo,
		)); err != nil {
			panic(err)
		}
	}

	sampleActor := module

	if err := fs.GenerateFile(atom.MakeServiceComponent(m.SourceDirectory, module, sampleActor)); err != nil {
		panic(err)
	}
}

func ServicePath() string {
	return "core/service"
}
