package generate

import (
	"fmt"

	"github.com/charmingruby/gentoo/internal/command/shared/constant"
	"github.com/charmingruby/gentoo/internal/command/shared/generator"
	"github.com/charmingruby/gentoo/internal/command/shared/validator"
	"github.com/ettle/strcase"
	"github.com/spf13/cobra"
)

const (
	DEFAULT_SERVICE_PKG = "service"
)

type generateServiceTemplateParams struct {
	PackageName string
	ServiceName string
}

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

			if err := generator.GenerateFile(input); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&name, "name", "n", "", "service name")
	cmd.Flags().StringVarP(&pkg, "pkg", "p", DEFAULT_SERVICE_PKG, "service package")

	return cmd
}

func (c *Command) makeServiceInput(module, resourceName, pkg string) generator.GenerateFileInput {
	sourceDir := c.config.BaseConfiguration.SourceDir

	camelCasePkgName := strcase.ToCamel(pkg)
	formattedResourceName := strcase.ToGoCase(resourceName, strcase.TitleCase, 0)

	// source_dir/module/core/service/name_service.go
	directory := fmt.Sprintf("%s/%s/core/%s/",
		sourceDir,
		module,
		camelCasePkgName,
	)

	return generator.GenerateFileInput{
		Module:       module,
		Resource:     "service",
		ResourceName: resourceName,
		Data: generateServiceTemplateParams{
			PackageName: camelCasePkgName,
			ServiceName: formattedResourceName,
		},
		Directory:  directory,
		Suffix:     "_" + pkg,
		ActionType: constant.GENERATE_ACTION,
		HasTest:    false,
	}
}

func (c *Command) validateServiceArgs(
	module string,
	resourceName string,
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
			FieldName: "pkg",
			Value:     pkg,
		},
	}

	if err := validator.ValidateArgsList(args); err != nil {
		return nil, err
	}

	return args, nil
}
