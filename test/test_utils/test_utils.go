package test_utils

import (
	"fmt"
	"log"
	"net/http/httptest"
	"os"
	"reflect"
	"sync"

	"github.com/joho/godotenv"
	"github.com/wesleyfebarretos/ticket-sale/config"
	"github.com/wesleyfebarretos/ticket-sale/infra/db"
	"github.com/wesleyfebarretos/ticket-sale/io/routes"
	"github.com/wesleyfebarretos/ticket-sale/middleware"
	"github.com/wesleyfebarretos/ticket-sale/test/test_container"
)

var runningContainers = []*test_container.ContainerResult{}

func BeforeAll() *httptest.Server {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("error on find working dir: %s", err.Error())
	}

	err = godotenv.Load(fmt.Sprintf("%s/.env.test", wd))
	if err != nil {
		log.Fatal("error loading .env file")
	}

	config.Init()

	pgContainer := test_container.SetupPG()
	runningContainers = append(runningContainers, pgContainer)
	config.Envs.DBPort = fmt.Sprintf("%d", pgContainer.Port)

	db.Init()

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
