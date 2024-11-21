package main

import (
	"github.com/charmingruby/gentoo/config"
	"github.com/charmingruby/gentoo/internal/command"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{}

	config, err := config.New("./dummy")
	if err != nil {
		panic(err)
	}

	cmd := command.New(rootCmd, &config)
	cmd.SetupGenerate()

	rootCmd.Execute()
}
