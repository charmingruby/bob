package generate

import (
	"fmt"

	"github.com/charmingruby/gentoo/internal/command/shared/constant"
	"github.com/charmingruby/gentoo/internal/command/shared/generator"
	"github.com/charmingruby/gentoo/internal/command/shared/validator"
	"github.com/spf13/cobra"
)

type generateHandlerTemplateParams struct {
	HandlerName string
}

func (c *Command) runGenerateHandler() *cobra.Command {
	var (
		module       string
		resourceName string
		variant      string
		pkg          string
	)

	cmd := &cobra.Command{
		Use:   "handler",
		Short: "Generates a new handler",
		Run: func(cmd *cobra.Command, args []string) {
			arguments, err := c.validateHandlerArgs(module, resourceName, variant, pkg)
			if err != nil {
				panic(err)
			}

			input := c.makeHandlerInput(
				arguments[0].CurrentState,
				arguments[1].CurrentState,
				arguments[2].CurrentState,
				arguments[3].CurrentState,
			)

			if err := generator.GenerateFile(input); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&resourceName, "resource name", "r", "", "handler name")
	cmd.Flags().StringVarP(&variant, "variant", "v", "", "comunication protocol")
	cmd.Flags().StringVarP(&pkg, "pkg", "p", "", "communication handler package")

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
	}
}

func (c *Command) validateHandlerArgs(
	module string,
	resourceName string,
	variant string,
	pkg string,
) ([]validator.Arg, error) {

	args := []validator.Arg{
		{
			FieldName:     "module",
			MustHaveState: true,
			CurrentState:  module,
			EmptyState:    "",
		},
		{
			FieldName:     "resourceName",
			MustHaveState: true,
			CurrentState:  resourceName,
			EmptyState:    "",
		},
		{
			FieldName:     "variant",
			MustHaveState: true,
			CurrentState:  variant,
			EmptyState:    "",
		},
		{
			FieldName:     "pkg",
			MustHaveState: true,
			CurrentState:  pkg,
			EmptyState:    "",
		},
	}

	if err := validator.ValidateArgsList(args); err != nil {
		return nil, err
	}

	return args, nil
}
