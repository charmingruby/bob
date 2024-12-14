package main

import (
	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/cli/command"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{}

	config, err := config.New()
	if err != nil {
		panic(err)
	}

	cmd := command.New(rootCmd, config)
	cmd.Setup()

	rootCmd.Execute()
}
