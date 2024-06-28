package migration

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/wesleyfebarretos/ticket-sale/config"
)

type MigrationLog struct {
	migrate.Logger
}

func (m *MigrationLog) Verbose() bool {
	return true
}

func (m *MigrationLog) Printf(format string, v ...interface{}) {
	logMessage := fmt.Sprintf(format, v...)
	migrationLog(logMessage)
}

type MigrationType string

const (
	MigrationTypeTable   MigrationType = "tables"
	MigrationTypeSeeders MigrationType = "seeders"
)

func Up() {
	pool, driver := openConnection()
	defer pool.Close()

	upMigration(createMigrationInstance(driver, MigrationTypeTable))

	upMigration(createMigrationInstance(driver, MigrationTypeSeeders))
}

func UpSeeders() {
	pool, driver := openConnection()
	defer pool.Close()

	upMigration(createMigrationInstance(driver, MigrationTypeSeeders))
}

func Down() {
	pool, driver := openConnection()
	defer pool.Close()

	downMigration(createMigrationInstance(driver, MigrationTypeTable))
}

func openConnection() (*pgxpool.Pool, database.Driver) {
	stringConn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBHost,
		config.Envs.DBPort,
		config.Envs.DBName,
	)
	poolConfig, err := pgxpool.ParseConfig(stringConn)
	if err != nil {
		log.Fatalf("error on parse pool config: %v", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalf("error on open pool connection: %v", err)
	}

	sqlDB := stdlib.OpenDB(*poolConfig.ConnConfig)

	driver, err := pgx.WithInstance(sqlDB, &pgx.Config{})
	if err != nil {
		log.Fatalf("error on create db instance: %v", err)
	}
	return pool, driver
}

func createMigrationInstance(driver database.Driver, migrationType MigrationType) *migrate.Migrate {
	migration, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://cmd/migrations/%s", migrationType),
		config.Envs.DBName,
		driver,
	)
	if err != nil {
		log.Fatalf("could not create migrate instance: %v", err)
	}

	migration.Log = &MigrationLog{}

	return migration
}

func upMigration(migration *migrate.Migrate) {
	err := migration.Up()

	if err != nil && fileNotFoundErr(err) {
		migrationLogWarning("files not found")
		err = nil
	}

	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("could not apply up migrations: %v", err)
	}
}

func downMigration(migration *migrate.Migrate) {
	err := migration.Down()

	if err != nil && fileNotFoundErr(err) {
		migrationLogWarning("not found files to tables migrations")
		err = nil
	}

	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("could not apply down migrations: %v", err)
	}
}

func fileNotFoundErr(err error) bool {
	return strings.Contains(err.Error(), "file does not exist")
}

func migrationLog(logMessage string) {
	log.Printf("migration [LOG]: %s", logMessage)
}

func migrationLogWarning(logMessage string) {
	log.Printf("migration [LOG_WARNING]: %s", logMessage)
}
