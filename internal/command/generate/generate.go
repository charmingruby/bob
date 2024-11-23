package generate

import (
	"github.com/charmingruby/bob/config"
	"github.com/spf13/cobra"
)

type Command struct {
	cmd    *cobra.Command
	config *config.Configuration
}

func New(cmd *cobra.Command, config *config.Configuration) *Command {
	return &Command{
		cmd:    cmd,
		config: config,
	}
}

func (c *Command) Setup() {
	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generates components",
	}

	generateCmd.AddCommand(c.runGenerateHandler())
	generateCmd.AddCommand(c.runGenerateModel())
	generateCmd.AddCommand(c.runGenerateService())

	c.cmd.AddCommand(generateCmd)
}
