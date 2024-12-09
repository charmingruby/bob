package postgres

import (
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/spf13/cobra"
)

func RunRepository(m filesystem.Manager) *cobra.Command {
	var (
		module           string
		modelName        string
		tableName        string
		needDependencies bool
	)

	cmd := &cobra.Command{
		Use:   "repository",
		Short: "Generates a new repository",
		Run: func(cmd *cobra.Command, args []string) {
			println("oi")
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&modelName, "model", "n", "", "model to be managed by the repository")
	cmd.Flags().StringVarP(&tableName, "table name", "t", "", "table name on migrations, by default is the model name")
	cmd.Flags().BoolVarP(&needDependencies, "dependencies", "d", false, "generate dependencies")

	return cmd
}
