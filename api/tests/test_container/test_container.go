package test_container

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
)

const PGPORT = 5432

type ContainerResult struct {
	container testcontainers.Container
	ctx       context.Context
	Host      string
	Port      uint
}

func (c ContainerResult) Terminate() {
	c.container.Terminate(c.ctx)
}

type TestLogConsumer struct{}

func (g *TestLogConsumer) Accept(l testcontainers.Log) {
	log.Print(string(l.Content))
}

func setupContainer(
	containerRequest testcontainers.ContainerRequest,
	nPort nat.Port,
	printContainerLogs bool,
) *ContainerResult {
	ctx := context.Background()
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: containerRequest,
		Started:          true,
	})
	if err != nil {
		log.Panicf("Failed to start container %+v, with error: %v", containerRequest, err)
	}

	host, err := container.Host(ctx)
	if err != nil {
		log.Panicf("Failed to retrieve host %+v, with error: %v", containerRequest, err)
	}

	fmt.Println("Requested port:", nPort)

	port, err := container.MappedPort(ctx, nPort)
	if err != nil {
		log.Panicf("Failed to retrieve port %+v, with error: %v", containerRequest, err)
	}

	if printContainerLogs {
		logConsumer := TestLogConsumer{}

		err = container.StartLogProducer(ctx)
		if err != nil {
			log.Panicf("%s", err)
		}
		container.FollowOutput(&logConsumer)
	}

	return &ContainerResult{
		container: container,
		ctx:       ctx,
		Host:      host,
		Port:      uint(port.Int()),
	}
}
