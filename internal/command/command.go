package command

import (
	"github.com/charmingruby/gentoo/config"
	"github.com/charmingruby/gentoo/internal/command/generate"
	"github.com/spf13/cobra"
)

type command struct {
	cmd    *cobra.Command
	config *config.Configuration
}

func New(cmd *cobra.Command, config *config.Configuration) *command {
	return &command{
		cmd:    cmd,
		config: config,
	}
}

func (c *command) SetupGenerate() {
	generate.New(c.cmd, c.config).Setup()
}
