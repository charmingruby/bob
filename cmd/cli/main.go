package main

import (
	"os"

	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/cli/command"
	"github.com/charmingruby/bob/internal/component/organism"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{}

	if _, err := os.Stat(config.CONFIG_FILE); os.IsNotExist(err) {
		organism.MakeAndRunConfigure()
	}

	config, err := config.New()
	if err != nil {
		panic(err)
	}

	if config != nil {
		command.New(rootCmd, *config).Setup()
	}

	rootCmd.Execute()
}
