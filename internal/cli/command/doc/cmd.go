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
		Short: "Presents architectural concepts",
		Long:  "This command presents the various architectural decisions for the project.",
	}

	cmd.AddCommand(
		RunArchicture(),
	)

	c.cmd.AddCommand(cmd)
}
