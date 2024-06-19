package test_container

import (
	"os"
	"os/exec"
)

func SetupPG() *ContainerResult {
	cmd := &exec.Cmd{}

	cmd = exec.Command("docker", "run", "--name", "ticket_sale", "--rm", "-d", "-e", "POSTGRES_DB=ticket_sale", "-e", "POSTGRES_PASSWORD=root", "-e", "POSTGRES_USER=root", "-p", "6432:5432", "postgres:13-alpine")

	cmd.Stdout = os.Stdout
	cmd.Run()

	// FIX:
	// This above works and below not
	// the code below are open container with sudo and my app cant access it
	// see how to resolve

	// req := testcontainers.ContainerRequest{
	// 	Image: "postgres:13-alpine",
	// 	Env: map[string]string{
	// 		"POSTGRES_PASSWORD": config.Envs.DBPassword,
	// 		"POSTGRES_DB":       config.Envs.DBName,
	// 		"POSTGRES_USER":     config.Envs.DBUser,
	// 	},
	// 	ExposedPorts: []string{
	// 		fmt.Sprintf("%s/tcp", config.Envs.DBPort),
	// 	},
	// 	WaitingFor: wait.ForLog("database system is ready to accept connections"),
	// 	Name:       "Ticket-sale-DB-IT",
	// }
	//
	// printContainerLogs := true
	//
	// return setupContainer(req, nat.Port(fmt.Sprintf("%s/tcp", config.Envs.DBPort)), printContainerLogs)
	return &ContainerResult{}
}
