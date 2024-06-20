package main

import (
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/wesleyfebarretos/ticket-sale/cmd/migrations/migration"
	"github.com/wesleyfebarretos/ticket-sale/config"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if len(os.Args) < 2 {
		log.Fatal("Action argument is required (up or down)")
	}

	config.Init()

	action := os.Args[1]
	if action == "up" {
		migration.Up()
	} else {
		migration.Down()
	}
	// var migrationType string

	//
	// dbConnection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
	// 	config.Envs.DBUser,
	// 	config.Envs.DBPassword,
	// 	config.Envs.DBHost,
	// 	config.Envs.DBPort,
	// 	config.Envs.DBName,
	// )
	//
	// migrationCommand := []string{"-database", dbConnection, "-path"}
	//
	// cmd := &exec.Cmd{}
	//
	// wd, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// switch migrationType {
	// case "tables":
	// 	tablesPath := wd + "/cmd/migrations/tables"
	// 	migrationCommand = append(migrationCommand, tablesPath)
	// case "seeders":
	// 	seedersPath := wd + "/cmd/migrations/seeders"
	// 	migrationCommand = append(migrationCommand, seedersPath)
	// default:
	// 	log.Fatal("Invalid migration type. Use 'tables' or 'seeders'")
	// }
	//
	// switch action {
	// case "up":
	// 	migrationCommand = append(migrationCommand, "up")
	// case "down":
	// 	migrationCommand = append(migrationCommand, "down")
	// default:
	// 	log.Fatal("Invalid action. Use 'up' or 'down'")
	// }
	//
	// cmd = exec.Command("migrate", migrationCommand...)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// input := bytes.Buffer{}
	// input.WriteString("Y\n")
	// cmd.Stdin = &input
	//
	// if err := cmd.Run(); err != nil {
	// 	log.Fatal(err.Error())
	// }
}
