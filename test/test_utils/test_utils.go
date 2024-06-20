package test_utils

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"sync"

	"github.com/joho/godotenv"
	"github.com/wesleyfebarretos/ticket-sale/cmd/app"
	"github.com/wesleyfebarretos/ticket-sale/cmd/migrations/migration"
	"github.com/wesleyfebarretos/ticket-sale/config"
	"github.com/wesleyfebarretos/ticket-sale/infra/db"
	"github.com/wesleyfebarretos/ticket-sale/test/test_container"
)

var runningContainers = []*test_container.ContainerResult{}

func BeforeAll() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("error on find working dir: %s", err.Error())
	}

	err = godotenv.Load(fmt.Sprintf("%s/.env.test", wd))
	if err != nil {
		log.Fatal("error loading .env file")
	}

	config.Init()

	wg := &sync.WaitGroup{}
	runInParallel(wg, func() {
		pgContainer := test_container.SetupPG()
		runningContainers = append(runningContainers, pgContainer)
		config.Envs.DBPort = fmt.Sprintf("%d", pgContainer.Port)
	})
	wg.Wait()

	db.Init()
	migration.Up()
	app.Run()
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

func printStruct(s interface{}) {
	v := reflect.ValueOf(s)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		key := t.Field(i)
		value := v.Field(i)
		fmt.Printf("Key: %v Value: %v\n", key.Name, value)
	}
}
