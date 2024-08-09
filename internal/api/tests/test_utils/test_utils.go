package test_utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"sync"
	"testing"

	"github.com/joho/godotenv"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/config"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/middleware"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/migrations"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/routes"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/tests/test_container"
	"github.com/wesleyfebarretos/ticket-sale/internal/infra/db"
)

var runningContainers = []*test_container.ContainerResult{}

func BeforeAll() *httptest.Server {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("error on find working dir: %s", err.Error())
	}

	if err := godotenv.Load(fmt.Sprintf("%s/.env.test", wd)); err != nil {
		log.Fatal("error loading .env file")
	}

	config.Init()

	// Set app enviroment to testing
	config.Envs.AppEnv = "testing"

	// Close error chanel, panics was spam the terminal making it difficult to tracking tests
	os.Stderr.Close()

	// Moving default logger output to Stderr which is closed to not spam the tests
	config.Envs.Logger.Output = os.Stderr

	pgContainer := test_container.SetupPG()
	runningContainers = append(runningContainers, pgContainer)
	config.Envs.DB.Port = fmt.Sprintf("%d", pgContainer.Port)

	db.Init()
	migrations.Up()

	server := httptest.NewServer(routes.Bind())
	config.Envs.PublicHost = fmt.Sprintf("http://localhost:%s", server.URL)

	return server
}

func Finish() {
	for _, container := range runningContainers {
		container.Terminate()
	}
}

func runInParallel(wg *sync.WaitGroup, work func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		work()
	}()
}

func GenerateJwtToken(role string) string {
	token, _, _ := middleware.JWT.TokenGenerator(&middleware.UserClaims{
		Id:   int32(1),
		Role: role,
	})

	return token
}

func PrintStruct(s interface{}) {
	v := reflect.ValueOf(s)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		key := t.Field(i)
		value := v.Field(i)
		fmt.Printf("Key: %v Value: %v\n", key.Name, value)
	}
}

func Decode[T any](t *testing.T, input io.Reader, into *T) {
	if err := json.NewDecoder(input).Decode(into); err != nil {
		t.Fatalf("could not parse response body: %v", err)
	}
}

func DebugResponse(body io.Reader) {
	b, _ := io.ReadAll(body)
	fmt.Println(string(b))
	os.Exit(1)
}

func DeleteRequiredField(t *testing.T, field string, targetMap map[string]any) {
	depthKeys := strings.Split(field, ".")
	if len(depthKeys) > 1 {
		mapEntry := targetMap
		for i := 0; i < len(depthKeys); i++ {
			if _, ok := mapEntry[depthKeys[i]]; !ok {
				t.Fatalf("key [%s] not found", depthKeys[i])
			}
			if i == len(depthKeys)-1 {
				delete(mapEntry, depthKeys[i])
				break
			}
			_, ok := mapEntry[depthKeys[i]].(map[string]any)
			if !ok {
				t.Fatalf("the key [%s] is not a map", depthKeys[i])
			}

			mapEntry = mapEntry[depthKeys[i]].(map[string]any)
		}
	} else {
		if _, ok := targetMap[depthKeys[0]]; !ok {
			t.Fatalf("key [%s] not found", depthKeys[0])
		}
		delete(targetMap, depthKeys[0])
	}
}
