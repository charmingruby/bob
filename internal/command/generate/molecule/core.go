package molecule

import (
	"github.com/charmingruby/bob/internal/command/generate/atom"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/input"
	"github.com/charmingruby/bob/internal/command/shared/fs"
	"github.com/spf13/cobra"
)

func RunCore(m component.Manager) *cobra.Command {
	var (
		module string
	)

	cmd := &cobra.Command{
		Use:   "core",
		Short: "Generates a new core molecule",
		Run: func(cmd *cobra.Command, args []string) {
			if err := input.ValidateOnlyModuleCommandInput(module); err != nil {
				panic(err)
			}

			MakeCore(m, module)
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")

	return cmd
}

func MakeCore(m component.Manager, module string) {
	if err := fs.GenerateDirectory(m.ModuleDirectory(module), "core"); err != nil {
		panic(err)
	}

	if err := fs.GenerateMultipleDirectories(
		component.ModulePath(m.SourceDirectory, module, CorePath()),
		[]string{"service", "model", "repository"},
	); err != nil {
		panic(err)
	}

	sampleActor := module

	MakeService(m, "", module)

	if err := fs.GenerateFile(atom.MakeRepositoryComponent(m, module, sampleActor)); err != nil {
		panic(err)
	}

	if err := fs.GenerateFile(atom.MakeModelComponent(m, module, sampleActor)); err != nil {
		panic(err)
	}
}

func CorePath() string {
	return "core"
}
