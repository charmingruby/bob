package command

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/cli/command/add"
	"github.com/charmingruby/bob/internal/cli/command/doc"
	"github.com/charmingruby/bob/internal/cli/command/setup"
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
	if !cfgFileExists {
		setup.Setup().Execute()

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
	configurationName := "Configuration"

	var commandChoice string
	prompt := &survey.Select{
		Message: input.ChooseSectionMessage(section),
		Options: []string{templateName, componentsName, documentationName, configurationName},
	}
	survey.AskOne(prompt, &commandChoice)

	switch commandChoice {
	case templateName:
		template.Setup(fs).Execute()
	case componentsName:
		add.Setup(fs).Execute()
	case documentationName:
		doc.Setup().Execute()
	case configurationName:
		setup.Setup().Execute()
	}

	os.Exit(0)
}
