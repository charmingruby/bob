package single

import (
	"fmt"

	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/command/shared/constant"
	"github.com/charmingruby/bob/internal/command/shared/fs"
	"github.com/charmingruby/bob/internal/command/shared/validator"
	"github.com/spf13/cobra"
)

const (
	HANDLER_TEMPLATE        = "handler"
	DEFAULT_HANDLER_VARIANT = "rest"
	DEFAULT_HANDLER_PKG     = "endpoint"
)

func RunHandler(cfg config.Configuration) *cobra.Command {
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
			arguments, err := validateHandlerArgs(module, name, variant, pkg)
			if err != nil {
				panic(err)
			}

			component := makeHandlerComponent(
				cfg.BaseConfiguration.RootDir,
				cfg.BaseConfiguration.SourceDir,
				arguments[0].Value,
				arguments[1].Value,
				arguments[2].Value,
				arguments[3].Value,
			)

			file := fs.File{
				CommandType:          constant.GENERATE_COMMAND,
				TemplateName:         HANDLER_TEMPLATE,
				TemplateData:         component.Data,
				FileName:             component.Name,
				FileSuffix:           component.Package.Name,
				ResourceName:         component.Name,
				ResourceSuffix:       component.Suffix,
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
	cmd.Flags().StringVarP(&variant, "variant", "v", DEFAULT_HANDLER_VARIANT, "comunication protocol")
	cmd.Flags().StringVarP(&pkg, "pkg", "p", DEFAULT_HANDLER_PKG, "communication handler package")

	return cmd
}

func makeHandlerComponent(rootDir, srcDir, module, name, variant, pkg string) Single {
	component := New(SingleInput{
		Module:      module,
		Name:        name,
		PackageName: pkg,
		Suffix:      pkg,
		HasTest:     false,
	}, WithDefaultTemplate())

	// source_dir/module/transport/protocol/handler_name/resource_handler.go
	directory := fmt.Sprintf("%s/%s/%s/transport/%s/%s",
		rootDir,
		srcDir,
		component.Module,
		variant,
		component.Package.Name,
	)

	component.Directory = directory

	return *component
}

func validateHandlerArgs(
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
