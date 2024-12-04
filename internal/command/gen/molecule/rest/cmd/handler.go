package cmd

import (
	"github.com/charmingruby/bob/internal/command/gen/molecule/rest/rest_component"
	"github.com/charmingruby/bob/internal/command/shared/component/input"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunHandler(m filesystem.Manager) *cobra.Command {
	var (
		module string
		name   string
	)

	cmd := &cobra.Command{
		Use:   "handler",
		Short: "Generates a new handler",
		Run: func(cmd *cobra.Command, args []string) {
			if err := input.ValidateDefaultCommandInput(module, name); err != nil {
				panic(err)
			}

			if err := m.GenerateFile(rest_component.MakeRequest(
				m,
				module,
				name,
			)); err != nil {
				panic(err)
			}

			if err := m.GenerateFile(rest_component.MakeResponse(
				m,
				module,
				name,
			)); err != nil {
				panic(err)
			}

			if err := m.GenerateFile(rest_component.MakeHandlerComponent(
				m,
				module,
				name,
			)); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&name, "name", "n", "", "handler name")

	return cmd
}
