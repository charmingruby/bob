package doc

import (
	"fmt"

	"github.com/charmingruby/bob/internal/shared/definition/architecture"
	"github.com/spf13/cobra"
)

func RunArchicture() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "arch",
		Short: "Presents the overall architecture",
		Long:  "This command presents the overall architecture of the project.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(architecture.PROPOSAL)
		},
	}

	return cmd
}
