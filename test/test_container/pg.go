package test_container

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/wesleyfebarretos/ticket-sale/config"
)

func SetupPG() *ContainerResult {
	ctx := context.Background()

	postgresContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:13-alpine"),
		postgres.WithDatabase(config.Envs.DBName),
		testcontainers.WithLogConsumers(&TestLogConsumer{}),
		postgres.WithUsername(config.Envs.DBUser),
		postgres.WithPassword(config.Envs.DBPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Fatal(err)
	}

	host, err := postgresContainer.Host(ctx)
	if err != nil {
		log.Fatalf("Host error: %s", err)
	}
	port, err := postgresContainer.MappedPort(ctx, nat.Port(fmt.Sprintf("%d/tcp", 5432)))
	if err != nil {
		log.Fatalf("Port error: %s", err)
	}

	return &ContainerResult{
		container: postgresContainer,
		ctx:       ctx,
		Host:      host,
		Port:      uint(port.Int()),
	}
}
