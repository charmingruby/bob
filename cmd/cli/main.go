package main

import (
	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/cli/command"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{}

	cfg, cfgFileExists, err := config.New()
	if err != nil {
		panic(err)
	}

	command.New(rootCmd, cfg).Setup(cfgFileExists)

	rootCmd.Execute()
}
