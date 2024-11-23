package generate

import (
	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/command/generate/single"
	"github.com/spf13/cobra"
)

type Command struct {
	cmd    *cobra.Command
	config config.Configuration
}

func New(cmd *cobra.Command, config config.Configuration) *Command {
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

	generateCmd.AddCommand(single.RunModel(c.config))
	generateCmd.AddCommand(single.RunService(c.config))
	generateCmd.AddCommand(single.RunHandler(c.config))

	c.cmd.AddCommand(generateCmd)
}
