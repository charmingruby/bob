package generate

import (
	"fmt"

	"github.com/charmingruby/bob/internal/command/shared/constant"
	"github.com/charmingruby/bob/internal/command/shared/generator"
	"github.com/charmingruby/bob/internal/command/shared/validator"
	"github.com/spf13/cobra"
)

const (
	DEFAULT_HANDLER_VARIANT = "rest"
	DEFAULT_HANDLER_PKG     = "endpoint"
)

type generateHandlerTemplateParams struct {
	HandlerName string
}

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

			if err := generator.GenerateFile(input); err != nil {
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

func (c *Command) makeHandlerInput(module, resourceName, variant, pkg string) generator.GenerateFileInput {
	sourceDir := c.config.BaseConfiguration.SourceDir

	// source_dir/module/transport/protocol/handler_name/resource_handler.go
	directory := fmt.Sprintf("%s/%s/transport/%s/%s/",
		sourceDir,
		module,
		variant,
		pkg,
	)

	return generator.GenerateFileInput{
		Module:       module,
		Resource:     "handler",
		ResourceName: resourceName,
		Data:         generateHandlerTemplateParams{HandlerName: resourceName},
		Directory:    directory,
		Suffix:       "_handler",
		ActionType:   constant.GENERATE_ACTION,
		HasTest:      false,
	}
}

func (c *Command) validateHandlerArgs(
	module string,
	resourceName string,
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
			Value:      resourceName,
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
