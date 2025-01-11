package component

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/definition/component/base"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/charmingruby/bob/tooling/script/migrator"
)

const (
	baseUpSQL = `CREATE TABLE IF NOT EXISTS %s
(
	id varchar PRIMARY KEY NOT NULL,
	name varchar NOT NULL,
	created_at timestamp DEFAULT now() NOT NULL,
	updated_at timestamp,
	deleted_at timestamp
);
`
	baseDownSQL = `DROP TABLE IF EXISTS %s;`
)

func RunMigration(m filesystem.Manager, tableName string) ([]filesystem.File, error) {
	tmpFile, err := os.CreateTemp("", "Makefile")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write([]byte(migrator.Makefile))
	if err != nil {
		return nil, err
	}
	err = tmpFile.Close()
	if err != nil {
		return nil, err
	}

	migrationNameParam := fmt.Sprintf("NAME=%s", tableName)

	e := exec.Command("make", "-f", tmpFile.Name(), "new-mig", migrationNameParam)
	var out bytes.Buffer
	e.Stdout = &out
	e.Stderr = &out

	err = e.Run()
	if err != nil {
		return nil, err
	}

	migrationsDir := m.MainDirectory() + definition.SQL_MIGRATION
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return nil, err
	}

	var upFile, downFile string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".up.sql") {
			upFile = filepath.Join(migrationsDir, file.Name())
		} else if strings.HasSuffix(file.Name(), ".down.sql") {
			downFile = filepath.Join(migrationsDir, file.Name())
		}
	}

	if upFile != "" {
		upSQL := fmt.Sprintf(baseUpSQL, tableName)
		err = os.WriteFile(upFile, []byte(upSQL), 0644)
		if err != nil {
			return nil, err
		}
	}

	if downFile != "" {
		downSQL := fmt.Sprintf(baseDownSQL, tableName)
		err = os.WriteFile(downFile, []byte(downSQL), 0644)
		if err != nil {
			return nil, err
		}
	}

	components := []filesystem.File{
		{
			Identifier: base.BuildNonModuleIdentifier(
				"migration",
				fmt.Sprintf("%s down migration", tableName),
				migrationsDir,
			),
		},
		{
			Identifier: base.BuildNonModuleIdentifier(
				"migration",
				fmt.Sprintf("%s up migration", tableName),
				migrationsDir,
			),
		},
	}

	return components, nil
}
