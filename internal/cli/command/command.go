package command

import (
	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/cli/command/add"
	"github.com/charmingruby/bob/internal/cli/command/create"
	"github.com/charmingruby/bob/internal/cli/command/initialize"

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

func (c *command) Setup(cfgFileExists bool) {
	initialize.New(c.cmd)

	if !cfgFileExists {
		return
	}

	create.New(c.cmd, c.config).Setup()
	add.New(c.cmd, c.config).Setup()
}
