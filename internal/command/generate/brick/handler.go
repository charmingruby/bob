package brick

import (
	"fmt"

	"github.com/charmingruby/bob/internal/command/shared/constant"
	"github.com/charmingruby/bob/internal/command/shared/fs"
	"github.com/charmingruby/bob/internal/command/shared/validator/input"
	"github.com/spf13/cobra"
)

const (
	DEFAULT_HANDLER_VARIANT = "rest"
	DEFAULT_HANDLER_PKG     = "endpoint"
)

func RunHandler(destinationDirectory string) *cobra.Command {
	var (
		module string
		name   string
	)

	cmd := &cobra.Command{
		Use:   "handler",
		Short: "Generates a new handler",
		Run: func(cmd *cobra.Command, args []string) {
			if err := input.ValidateDefaultCommandArgs(module, name); err != nil {
				panic(err)
			}

			component := MakeHandlerComponent(
				destinationDirectory,
				module,
				name,
			)

			file := fs.File{
				CommandType:          constant.GENERATE_COMMAND,
				TemplateName:         constant.HANDLER_TEMPLATE,
				TemplateData:         component.Data,
				FileName:             component.Name,
				FileSuffix:           "handler",
				ResourceName:         component.Name,
				DestinationDirectory: component.Directory,
				HasTest:              component.HasTest,
			}

			if err := fs.GenerateFile(file); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&name, "name", "n", "", "handler name")

	return cmd
}

func MakeHandlerComponent(destinationDirectory, module, name string) Component {
	component := New(ComponentInput{
		Module:  module,
		Name:    name,
		Suffix:  "handler",
		HasTest: false,
	}, WithDefaultTemplate())

	component.Directory = fmt.Sprintf("%s/%s/transport/rest/endpoint",
		destinationDirectory,
		component.Module,
	)

	return *component
}
