package doc

import (
	"github.com/spf13/cobra"
)

type Command struct {
	cmd *cobra.Command
}

func New(cmd *cobra.Command) *Command {
	return &Command{
		cmd: cmd,
	}
}

func (c *Command) Setup() {
	cmd := &cobra.Command{
		Use:   "doc",
		Short: "Generates components",
	}

	cmd.AddCommand(
		RunArchicture(),
	)

	c.cmd.AddCommand(cmd)
}
