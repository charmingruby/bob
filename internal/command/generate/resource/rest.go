package single

import (
	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/fs"
	"github.com/charmingruby/bob/internal/command/shared/validator"
	"github.com/spf13/cobra"
)

const (
	REST_IDENTIFIER           = "rest"
	DEFAULT_REST_HANDLER_NAME = "handler"
	DEFAULT_REST_PKG          = "endpoint"
)

func RunRest(cfg config.Configuration) *cobra.Command {
	var (
		module  string
		pkg     string
		handler string
	)

	cmd := &cobra.Command{
		Use:   "rest",
		Short: "Generates a new rest resource",
		Run: func(cmd *cobra.Command, args []string) {
			arguments, err := validateRestArgs(module, pkg, handler)
			if err != nil {
				panic(err)
			}

			input := makeRestComponent(
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
	cmd.Flags().StringVarP(&handler, "handler name", "n", DEFAULT_REST_HANDLER_NAME, "http communication handler name")
	cmd.Flags().StringVarP(&pkg, "pkg", "p", DEFAULT_REST_PKG, "communication rest package")

	return cmd
}

func makeRestComponent(rootDir, srcDir, module, name, pkg string) component.Single {
	return *component.New(component.SingleInput{})
}

func validateRestArgs(
	module string,
	handler string,
	pkg string,
) ([]*validator.Arg, error) {
	args := []*validator.Arg{
		{
			FieldName:  "module",
			Value:      module,
			IsRequired: true,
		},
		{
			FieldName: "handler name",
			Value:     handler,
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
