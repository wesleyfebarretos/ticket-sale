package migration

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/wesleyfebarretos/ticket-sale/config"
	"github.com/wesleyfebarretos/ticket-sale/utils"
)

func Up() {
	// TODO: IMPROVE FOLDER ORGANIZATION AND RENAME METHODS
	dbConnection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBHost,
		config.Envs.DBPort,
		config.Envs.DBName,
	)

	wd := utils.GetRootDir()
	tablesPath := fmt.Sprintf("%s/cmd/migrations/tables", wd)

	cmdTablesUp := exec.Command(
		"migrate",
		"-database",
		dbConnection,
		"-path",
		tablesPath,
		"up",
	)

	cmdd := exec.Command("ls", "-a")
	cmdd.Stdout = os.Stdout
	cmdd.Stderr = os.Stderr

	cmdTablesUp.Stdout = os.Stdout
	cmdTablesUp.Stderr = os.Stderr
	fmt.Println(cmdTablesUp)

	if err := cmdTablesUp.Run(); err != nil {
		log.Fatalf("Migration error: %s", err.Error())
	}
}
