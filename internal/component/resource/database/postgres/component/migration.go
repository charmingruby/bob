package component

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/charmingruby/bob/internal/component/resource/database/postgres/constant"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/charmingruby/bob/tooling/makefile"
)

const (
	baseUpSQL = `CREATE TABLE IF NOT EXISTS %s
(
	id varchar PRIMARY KEY NOT NULL,
	name varchar NOT NULL,
	created_at timestamp DEFAULT now() NOT NULL,
	deleted_at timestamp
);
`
	baseDownSQL = `DROP TABLE IF EXISTS %s;`
)

func RunMigration(m filesystem.Manager, tableName string) {
	tmpFile, err := os.CreateTemp("", "Makefile")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write([]byte(makefile.Makefile))
	if err != nil {
		log.Fatal(err)
	}
	err = tmpFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	migrationNameParam := fmt.Sprintf("NAME=%s", tableName)

	e := exec.Command("make", "-f", tmpFile.Name(), "new-mig", migrationNameParam)
	var out bytes.Buffer
	e.Stdout = &out
	e.Stderr = &out

	err = e.Run()
	if err != nil {
		log.Fatal(err)
	}

	migrationsDir := m.MainDirectory() + constant.MIGRATIONS_DIR
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		log.Fatal(err)
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
			log.Fatal(err)
		}
	}

	if downFile != "" {
		downSQL := fmt.Sprintf(baseDownSQL, tableName)
		err = os.WriteFile(downFile, []byte(downSQL), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}
