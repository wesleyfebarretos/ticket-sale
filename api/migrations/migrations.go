package migrations

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
	"github.com/wesleyfebarretos/ticket-sale/api/config"
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
	MigrationTypeTable     MigrationType = "tables"
	MigrationTypeSeeders   MigrationType = "seeders"
	MigrationTypeViews     MigrationType = "views"
	MigrationTypeSchema    MigrationType = "schemas"
	tablesMigrationsTable  string        = "tables_migrations"
	seedersMigrationsTable string        = "seeders_migrations"
	viewsMigrationsTable   string        = "views_migrations"
	schemasMigrationsTable string        = "schemas_migrations"
)

func Up() {
	UpSchemas()
	UpTables()
	UpViews()
	UpSeeders(true)
}

func Down() {
	pool, driver := openConnection(tablesMigrationsTable)
	pool2, driver2 := openConnection(seedersMigrationsTable)
	pool3, driver3 := openConnection(viewsMigrationsTable)
	pool4, driver4 := openConnection(schemasMigrationsTable)
	defer pool.Close()
	defer pool2.Close()
	defer pool3.Close()
	defer pool4.Close()

	downMigration(createMigrationInstance(driver2, MigrationTypeSeeders, true))
	downMigration(createMigrationInstance(driver3, MigrationTypeViews, true))
	downMigration(createMigrationInstance(driver, MigrationTypeTable, true))
	downMigration(createMigrationInstance(driver4, MigrationTypeSchema, true))
}

func UpTables() {
	pool, driver := openConnection(tablesMigrationsTable)
	defer pool.Close()

	upMigration(createMigrationInstance(driver, MigrationTypeTable, true))
}

func UpSeeders(activeLogger bool) {
	pool, driver := openConnection(seedersMigrationsTable)
	defer pool.Close()

	upMigration(createMigrationInstance(driver, MigrationTypeSeeders, activeLogger))
}

func UpViews() {
	pool, driver := openConnection(viewsMigrationsTable)
	defer pool.Close()

	upMigration(createMigrationInstance(driver, MigrationTypeViews, true))
}

func UpSchemas() {
	pool, driver := openConnection(schemasMigrationsTable)
	defer pool.Close()

	upMigration(createMigrationInstance(driver, MigrationTypeSchema, true))
}

func openConnection(migrationsTable string) (*pgxpool.Pool, database.Driver) {
	stringConn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.Envs.DB.User,
		config.Envs.DB.Password,
		config.Envs.DB.Host,
		config.Envs.DB.Port,
		config.Envs.DB.Name,
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

	driver, err := pgx.WithInstance(sqlDB, &pgx.Config{
		MigrationsTable: migrationsTable,
	})
	if err != nil {
		log.Fatalf("error on create db instance: %v", err)
	}
	return pool, driver
}

func createMigrationInstance(driver database.Driver, migrationType MigrationType, activeLogger bool) *migrate.Migrate {
	migration, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://api/migrations/%s", migrationType),
		config.Envs.DB.Name,
		driver,
	)
	if err != nil {
		log.Fatalf("could not create migrate instance: %v", err)
	}

	if activeLogger {
		migration.Log = &MigrationLog{}
	}

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
