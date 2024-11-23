package single

import (
	"fmt"

	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/constant"
	"github.com/charmingruby/bob/internal/command/shared/fs"
	"github.com/charmingruby/bob/internal/command/shared/validator"
	"github.com/spf13/cobra"
)

const (
	MODEL_IDENTIFIER = "model"

	DEFAULT_MODEL_PKG = "model"
)

func RunModel(cfg config.Configuration) *cobra.Command {
	var (
		module string
		name   string
		pkg    string
	)

	cmd := &cobra.Command{
		Use:   "model",
		Short: "Generates a new model",
		Run: func(cmd *cobra.Command, args []string) {
			arguments, err := validateModelArgs(module, name, pkg)
			if err != nil {
				panic(err)
			}

			input := makeModelComponent(
				cfg.BaseConfiguration.RootDir,
				cfg.BaseConfiguration.SourceDir,
				arguments[0].Value,
				arguments[1].Value,
				arguments[2].Value,
			)

			if err := fs.GenerateFile(input); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&name, "name", "n", "", "model name")
	cmd.Flags().StringVarP(&pkg, "pkg", "p", DEFAULT_MODEL_PKG, "model package")

	return cmd
}

func makeModelComponent(rootDir, srcDir, module, name, pkg string) component.Single {
	component := component.New(component.SingleInput{
		Identifier:  MODEL_IDENTIFIER,
		ActionType:  constant.GENERATE_ACTION,
		Module:      module,
		Name:        name,
		PackageName: pkg,
		HasTest:     true,
	}, component.WithDefaultTemplateParams())

	// source_dir/module/core/pkg_name/model_name.go
	// source_dir/module/core/pkg_name/model_name_test.go
	directory := fmt.Sprintf("%s/%s/%s/core/%s",
		rootDir,
		srcDir,
		module,
		component.Package.Name,
	)

	component.Directory = directory

	return *component
}

func validateModelArgs(
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
