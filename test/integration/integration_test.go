package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/wesleyfebarretos/ticket-sale/internal/enum"
	_ "github.com/wesleyfebarretos/ticket-sale/test/test_init"
	"github.com/wesleyfebarretos/ticket-sale/test/test_utils"
)

var (
	Server            = &httptest.Server{}
	UserClientRequest = &http.Client{}
)

func TestMain(m *testing.M) {
	log.Println("Starting test setup")
	start := time.Now()

	// Setup containers, db, migrations and server.
	Server = test_utils.BeforeAll()

	serverUrl, err := url.Parse(Server.URL)
	if err != nil {
		setupError("could not parse server url: ", err)
	}

	// Make a new http client with user JWT.
	UserClientRequest = test_utils.NewHTTPClient(serverUrl, enum.USER_ROLE)

	log.Printf("Setup took %s seconds\n", time.Since(start))

	exitVal := m.Run()

	test_utils.Finish()

	os.Exit(exitVal)
}

func setupError(format string, v ...interface{}) {
	log.Fatalf("integration test [SETUP_ERROR]: "+format, v...)
}

func TError(t *testing.T, format string, v ...interface{}) {
	t.Errorf("integration test [FAIL_ERROR]: "+format, v...)
	log.Println()
}

func TErrorFatal(t *testing.T, format string, v ...interface{}) {
	t.Fatalf("integration test [FAIL_FATAL_ERROR]: "+format, v...)
	log.Println()
}

func TMakeRequest(t *testing.T, method, endpoint string, data any) *http.Response {
	url := fmt.Sprintf("%s/%s", Server.URL, endpoint)

	jsonData, err := json.Marshal(data)
	if err != nil {
		TErrorFatal(t, "could not marshal response body: %v", err)
	}

	body := bytes.NewReader(jsonData)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		TErrorFatal(t, "could not open a new request to path: %s", endpoint)
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := UserClientRequest.Do(req)
	if err != nil {
		TErrorFatal(t, "request failed %v", err)
	}

	return res
}

func TPointer[T any](value T) *T {
	return &value
}

func TDecode[T any](t *testing.T, input io.Reader, into *T) {
	if err := json.NewDecoder(input).Decode(&into); err != nil {
		TErrorFatal(t, "failed to decode response body: %v", err)
	}
}
