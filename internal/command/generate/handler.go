package generate

import (
	"fmt"

	"github.com/charmingruby/gentoo/internal/command/shared/generator"
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
			input := c.makeHandlerInput(module, resourceName, variant, pkg)

			err := generator.GenerateFile(input)
			if err != nil {
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
	}
}
