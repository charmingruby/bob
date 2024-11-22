package generate

import (
	"fmt"

	"github.com/charmingruby/gentoo/internal/command/shared/constant"
	"github.com/charmingruby/gentoo/internal/command/shared/generator"
	"github.com/charmingruby/gentoo/internal/command/shared/validator"
	"github.com/ettle/strcase"
	"github.com/spf13/cobra"
)

type generateModelTemplateParams struct {
	ModelName string
}

func (c *Command) runGenerateModel() *cobra.Command {
	var (
		module string
		name   string
		pkg    string
	)

	cmd := &cobra.Command{
		Use:   "model",
		Short: "Generates a new model",
		Run: func(cmd *cobra.Command, args []string) {
			arguments, err := c.validateModelArgs(module, name, pkg)
			if err != nil {
				panic(err)
			}

			input := c.makeModelInput(
				arguments[0].CurrentState,
				arguments[1].CurrentState,
				arguments[2].CurrentState,
			)

			if err := generator.GenerateFile(input); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&name, "name", "n", "", "model name")
	cmd.Flags().StringVarP(&pkg, "pkg", "p", "", "model package")

	return cmd
}

func (c *Command) makeModelInput(module, resourceName, pkg string) generator.GenerateFileInput {
	sourceDir := c.config.BaseConfiguration.SourceDir

	// source_dir/module/core/pkg_name/model_name.go
	// source_dir/module/core/pkg_name/model_name_test.go
	directory := fmt.Sprintf("%s/%s/core/%s",
		sourceDir,
		module,
		pkg,
	)

	formattedResourceName := strcase.ToGoCase(resourceName, strcase.TitleCase, 0)

	return generator.GenerateFileInput{
		Module:       module,
		Resource:     "model",
		ResourceName: resourceName,
		Data:         generateModelTemplateParams{ModelName: formattedResourceName},
		Directory:    directory,
		Suffix:       "",
		ActionType:   constant.GENERATE_ACTION,
		HasTest:      true,
	}
}

func (c *Command) validateModelArgs(
	module string,
	name string,
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
			FieldName:     "name",
			MustHaveState: true,
			CurrentState:  name,
			EmptyState:    "",
		},
		{
			FieldName:     "pkg",
			MustHaveState: true,
			CurrentState:  pkg,
			EmptyState:    "",
			DefaultState:  "model",
		},
	}

	if err := validator.ValidateArgsList(args); err != nil {
		return nil, err
	}

	return args, nil
}
