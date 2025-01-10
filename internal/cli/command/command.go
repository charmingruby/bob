package command

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/cli/command/add"
	"github.com/charmingruby/bob/internal/cli/command/doc"
	"github.com/charmingruby/bob/internal/cli/command/initialize"
	"github.com/charmingruby/bob/internal/cli/command/template"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/shared/filesystem"

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
	c.runInteractive()
}

func (c *command) runInteractive() {
	fs := filesystem.New(c.config)

	section := "command type"

	templateName := "Templates"
	componentsName := "Components"
	documentationName := "Documentations"

	var commandChoice string
	prompt := &survey.Select{
		Message: input.ChooseSectionMessage(section),
		Options: []string{templateName, componentsName, documentationName},
	}
	survey.AskOne(prompt, &commandChoice)

	switch commandChoice {
	case templateName:
		template.Setup(fs).Execute()
	case componentsName:
		add.Setup(fs).Execute()
	case documentationName:
		doc.Setup().Execute()
	}

	os.Exit(0)
}
