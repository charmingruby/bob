package generate

import (
	"fmt"

	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/constant"
	"github.com/charmingruby/bob/internal/command/shared/gen"
	"github.com/charmingruby/bob/internal/command/shared/validator"
	"github.com/spf13/cobra"
)

const (
	SERVICE_IDENTIFIER = "service"

	DEFAULT_SERVICE_PKG = "service"
)

func (c *Command) runGenerateService() *cobra.Command {
	var (
		module string
		name   string
		pkg    string
	)

	cmd := &cobra.Command{
		Use:   "service",
		Short: "Generates a new service",
		Run: func(cmd *cobra.Command, args []string) {
			arguments, err := c.validateServiceArgs(module, name, pkg)
			if err != nil {
				panic(err)
			}

			input := c.makeServiceInput(
				arguments[0].Value,
				arguments[1].Value,
				arguments[2].Value,
			)

			if err := gen.GenerateFile(input); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&name, "name", "n", "", "service name")
	cmd.Flags().StringVarP(&pkg, "pkg", "p", DEFAULT_SERVICE_PKG, "service package")

	return cmd
}

func (c *Command) makeServiceInput(module, name, pkg string) component.Component {
	component := component.New(component.ComponentInput{
		Identifier:  SERVICE_IDENTIFIER,
		ActionType:  constant.GENERATE_ACTION,
		Module:      module,
		Name:        name,
		PackageName: pkg,
		Suffix:      pkg,
		HasTest:     false,
	}, component.WithDefaultTemplateParams())

	// source_dir/module/core/service/name_service.go
	directory := fmt.Sprintf("%s/%s/%s/core/%s",
		c.config.BaseConfiguration.RootDir,
		c.config.BaseConfiguration.SourceDir,
		module,
		component.Package.Name,
	)

	component.Directory = directory

	return *component
}

func (c *Command) validateServiceArgs(
	module string,
	name string,
	pkg string,
) ([]*validator.Arg, error) {
	args := []*validator.Arg{
		{
			FieldName:  "module",
			Value:      module,
			IsRequired: true,
		},
		{
			FieldName:  "name",
			Value:      name,
			IsRequired: true,
		},
		{
			FieldName: "pkg",
			Value:     pkg,
		},
	}

	if err := validator.ValidateArgsList(args); err != nil {
		return nil, err
	}

	return args, nil
}
