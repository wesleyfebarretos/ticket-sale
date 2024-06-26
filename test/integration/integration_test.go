package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/wesleyfebarretos/ticket-sale/config"
	"github.com/wesleyfebarretos/ticket-sale/infra/db"
	"github.com/wesleyfebarretos/ticket-sale/middleware"
	"github.com/wesleyfebarretos/ticket-sale/repository/sqlc"
	_ "github.com/wesleyfebarretos/ticket-sale/test/test_init"
	"github.com/wesleyfebarretos/ticket-sale/test/test_utils"
)

var (
	Server        = &httptest.Server{}
	ClientRequest = &http.Client{}
)

func TestMain(m *testing.M) {
	log.Println("Starting test setup")
	start := time.Now()

	// Setup containers, db, migrations and server.
	Server = test_utils.BeforeAll()

	log.Printf("Setup took %s seconds\n", time.Since(start))

	exitVal := m.Run()

	test_utils.Finish()

	os.Exit(exitVal)
}

func TMakeRequest(t *testing.T, method, endpoint string, data any) *http.Response {
	url := fmt.Sprintf("%s/%s", Server.URL, endpoint)

	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("could not marshal response body: %v", err)
	}

	body := bytes.NewReader(jsonData)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fatalf("could not open a new request to path: %s", endpoint)
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := ClientRequest.Do(req)
	if err != nil {
		t.Fatalf("request failed %v", err)
	}

	if res.StatusCode == http.StatusUnauthorized {
		t.Fatalf("unauthorized with status code: %v", res.StatusCode)
	}

	if res.StatusCode == http.StatusForbidden {
		t.Fatalf("unauthorized: status code: %v", res.StatusCode)
	}

	return res
}

func TPointer[T any](value T) *T {
	return &value
}

func TRun(testFunc func(*testing.T)) func(*testing.T) {
	return func(t *testing.T) {
		beforeEach()
		testFunc(t)
	}
}

func TSetCookieWithUser(t *testing.T, user sqlc.GetUserWithPasswordByEmailRow) {
	token, _, _ := middleware.JWT.TokenGenerator(user)

	cookie := &http.Cookie{
		Name:  config.Envs.CookieName,
		Value: token,
	}

	serverUrl, err := url.Parse(Server.URL)
	if err != nil {
		t.Errorf("could not parse server url: %v", err)
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Errorf("could not create cookie: %v", err)
	}
	ClientRequest.Jar = jar
	ClientRequest.Jar.SetCookies(serverUrl, []*http.Cookie{cookie})
}

func TSetCookie(t *testing.T, role string) {
	cookie := &http.Cookie{
		Name:  config.Envs.CookieName,
		Value: test_utils.GenerateJwtToken(role),
	}

	serverUrl, err := url.Parse(Server.URL)
	if err != nil {
		t.Errorf("could not parse server url: %v", err)
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Errorf("could not create cookie: %v", err)
	}
	ClientRequest.Jar = jar
	ClientRequest.Jar.SetCookies(serverUrl, []*http.Cookie{cookie})
}

func beforeEach() {
	ClientRequest.Jar = nil
	db.TruncateAll()
}

func fileNotFoundErr(err error) bool {
	return strings.Contains(err.Error(), "file does not exist")
}

func setupError(format string, v ...interface{}) {
	log.Fatalf("integration test [SETUP_ERROR]: "+format, v...)
}
