package integration_test

import (
	"log"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	_ "github.com/wesleyfebarretos/ticket-sale/test/test_init"
	"github.com/wesleyfebarretos/ticket-sale/test/test_utils"
)

var Server = &httptest.Server{}

func TestMain(m *testing.M) {
	log.Println("Starting test setup")
	start := time.Now()

	Server = test_utils.BeforeAll()
	// TODO:
	// Need to Create a new client
	// I already made this function in test_utils keep development
	log.Printf("Setup took %s seconds\n", time.Since(start))

	exitVal := m.Run()

	test_utils.Finish()

	os.Exit(exitVal)
}
