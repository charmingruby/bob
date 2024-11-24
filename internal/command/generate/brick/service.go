package brick

import (
	"fmt"

	"github.com/charmingruby/bob/internal/command/shared/constant"
	"github.com/charmingruby/bob/internal/command/shared/fs"
	"github.com/charmingruby/bob/internal/command/shared/validator/input"
	"github.com/spf13/cobra"
)

const (
	SERVICE_TEMPLATE = "service"

	DEFAULT_SERVICE_PKG = "service"
)

func RunService(destinationDirectory string) *cobra.Command {
	var (
		module string
		name   string
	)

	cmd := &cobra.Command{
		Use:   "service",
		Short: "Generates a new service",
		Run: func(cmd *cobra.Command, args []string) {
			if err := input.ValidateDefaultCommandArgs(module, name); err != nil {
				panic(err)
			}

			component := makeServiceComponent(
				destinationDirectory,
				module,
				name,
			)

			file := fs.File{
				CommandType:          constant.GENERATE_COMMAND,
				TemplateName:         constant.SERVICE_TEMPLATE,
				TemplateData:         component.Data,
				FileName:             component.Name,
				FileSuffix:           "service",
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
	cmd.Flags().StringVarP(&name, "name", "n", "", "service name")

	return cmd
}

func makeServiceComponent(destinationDirectory, module, name string) Component {
	component := New(ComponentInput{
		Module:  module,
		Name:    name,
		Suffix:  "service",
		HasTest: false,
	}, WithDefaultTemplate())

	component.Directory = fmt.Sprintf("%s/%s/core/service",
		destinationDirectory,
		module,
	)

	return *component
}
