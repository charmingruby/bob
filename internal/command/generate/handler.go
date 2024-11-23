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
	HANDLER_IDENTIFIER      = "handler"
	DEFAULT_HANDLER_VARIANT = "rest"
	DEFAULT_HANDLER_PKG     = "endpoint"
)

func (c *Command) runGenerateHandler() *cobra.Command {
	var (
		module  string
		name    string
		variant string
		pkg     string
	)

	cmd := &cobra.Command{
		Use:   "handler",
		Short: "Generates a new handler",
		Run: func(cmd *cobra.Command, args []string) {
			arguments, err := c.validateHandlerArgs(module, name, variant, pkg)
			if err != nil {
				panic(err)
			}

			input := c.makeHandlerInput(
				arguments[0].Value,
				arguments[1].Value,
				arguments[2].Value,
				arguments[3].Value,
			)

			if err := gen.GenerateFile(input); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&name, "name", "n", "", "handler name")
	cmd.Flags().StringVarP(&variant, "variant", "v", DEFAULT_HANDLER_VARIANT, "comunication protocol")
	cmd.Flags().StringVarP(&pkg, "pkg", "p", DEFAULT_HANDLER_PKG, "communication handler package")

	return cmd
}

func (c *Command) makeHandlerInput(module, name, variant, pkg string) component.Component {
	component := component.New(component.ComponentInput{
		ActionType:  constant.GENERATE_ACTION,
		Module:      module,
		Identifier:  HANDLER_IDENTIFIER,
		Name:        name,
		PackageName: pkg,
		Suffix:      pkg,
		HasTest:     false,
	}, component.WithDefaultTemplateParams())

	// source_dir/module/transport/protocol/handler_name/resource_handler.go
	directory := fmt.Sprintf("%s/%s/%s/transport/%s/%s",
		c.config.BaseConfiguration.RootDir,
		c.config.BaseConfiguration.SourceDir,
		component.Module,
		variant,
		component.Package.Name,
	)

	component.Directory = directory

	return *component
}

func (c *Command) validateHandlerArgs(
	module string,
	name string,
	variant string,
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
			FieldName: "variant",
			Value:     variant,
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
