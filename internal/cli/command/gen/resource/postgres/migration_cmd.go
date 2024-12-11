package postgres

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/resource"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/spf13/cobra"
)

func RunMigration(m filesystem.Manager) *cobra.Command {
	var (
		tableName string
	)

	cmd := &cobra.Command{
		Use:   "migration",
		Short: "Generates a new migration",
		Run: func(cmd *cobra.Command, args []string) {
			if err := ValidateMigrationCommandInput(tableName); err != nil {
				panic(err)
			}

			resource.MakeAndRunPostgresMigration(m, tableName)
		},
	}

	cmd.Flags().StringVarP(&tableName, "table name", "t", "", "table name on migrations, by default, if it is not set, it will be not created")

	return cmd
}

func ValidateMigrationCommandInput(tableName string) error {
	args := []input.Arg{
		{
			FieldName:  "table name",
			Value:      tableName,
			IsRequired: true,
		},
	}

	if err := input.ValidateArgsList(args); err != nil {
		return err
	}

	return nil
}
