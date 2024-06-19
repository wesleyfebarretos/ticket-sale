package integration_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/wesleyfebarretos/ticket-sale/test/test_utils"
)

func TestGetAllUsers(t *testing.T) {}

func TestMain(m *testing.M) {
	log.Println("Starting test setup")
	start := time.Now()

	test_utils.BeforeAll()
	log.Printf("Setup took %s seconds\n", time.Since(start))

	exitVal := m.Run()

	test_utils.Finish()

	os.Exit(exitVal)
}
