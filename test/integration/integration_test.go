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
		ErrorFatal("could not parse server url")
	}

	// Make a new http client with user JWT.
	UserClientRequest = test_utils.NewHTTPClient(serverUrl, enum.USER_ROLE)

	log.Printf("Setup took %s seconds\n", time.Since(start))

	exitVal := m.Run()

	test_utils.Finish()

	os.Exit(exitVal)
}

func ErrorFatal(msg any) {
	switch v := msg.(type) {
	case string:
		log.Fatalf("integration test [ERROR_FATAL]: %s", v)
	case error:
		log.Fatalf("integration test [ERROR_FATAL]: %v", v)
	case int, int8, int16, int32, int64:
		log.Fatalf("integration test [ERROR_FATAL]: %d", v)
	case uint, uint8, uint16, uint32, uint64:
		log.Fatalf("integration test [ERROR_FATAL]: %d", v)
	case float32, float64:
		log.Fatalf("integration test [ERROR_FATAL]: %f", v)
	case bool:
		log.Fatalf("integration test [ERROR_FATAL]: %t", v)
	default:
		log.Fatalf("integration test [ERROR_FATAL]: %v", v)
	}
}

func TError(t *testing.T, format string, v ...interface{}) {
	t.Errorf("\nintegration test [FAIL_ERROR]: "+format, v...)
}

func TErrorFatal(t *testing.T, format string, v ...interface{}) {
	t.Fatalf("\nintegration test [FAIL_FATAL_ERROR]: "+format, v...)
}

func TMakeRequest(t *testing.T, method, endpoint string, data any) *http.Response {
	url := fmt.Sprintf("%s/%s", Server.URL, endpoint)

	jsonData, err := json.Marshal(data)
	if err != nil {
		ErrorFatal(err)
	}

	body := bytes.NewReader(jsonData)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		ErrorFatal("could not open a new request to path: " + endpoint)
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
