package generate

import (
	"github.com/charmingruby/gentoo/config"
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
		Short: "Generates resources",
	}

	generateCmd.AddCommand(c.runGenerateHandler())
	generateCmd.AddCommand(c.runGenerateModel())

	c.cmd.AddCommand(generateCmd)
}
