package resource

import (
	"fmt"

	"github.com/charmingruby/bob/internal/command/generate/brick"
	"github.com/charmingruby/bob/internal/command/shared/fs"
	"github.com/charmingruby/bob/internal/command/shared/validator/input"
	"github.com/spf13/cobra"
)

func RunRest(projectData, destinationDirectory string) *cobra.Command {
	var (
		module string
	)

	cmd := &cobra.Command{
		Use:   "rest",
		Short: "Generates a new rest resource",
		Run: func(cmd *cobra.Command, args []string) {
			if err := input.ValidateOnlyModuleCommandInput(module); err != nil {
				panic(err)
			}

			moduleDir := fmt.Sprintf("%s/%s", destinationDirectory, module)

			if err := fs.GenerateNestedDirectories(
				moduleDir,
				[]string{"transport", "rest", "endpoint"},
			); err != nil {
				panic(err)
			}

			if err := fs.GenerateFile(makeRestEntryBrickComponent(
				fmt.Sprintf("%s/%s", moduleDir, "transport/rest/endpoint"),
				fmt.Sprintf("%s/%s", projectData, destinationDirectory),
				module,
			)); err != nil {
				panic(err)
			}

			if err := fs.GenerateFile(brick.MakeHandlerComponent(
				destinationDirectory,
				module,
				"ping",
			)); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")

	return cmd
}

type restEntryBrickData struct {
	Module     string
	SourcePath string
}

func makeRestEntryBrickComponent(destinationDirectory, sourcePath, module string) fs.File {
	return makeEntryBrick(entryBrickParams{
		Module:       module,
		TemplateName: "rest_entry",
		TemplateData: restEntryBrickData{
			Module:     module,
			SourcePath: sourcePath,
		},
		EntryName:            "endpoint",
		DestinationDirectory: destinationDirectory,
	})
}
