package test_utils

import (
	"fmt"
	"log"
	"reflect"
	"sync"

	"github.com/joho/godotenv"
	"github.com/wesleyfebarretos/ticket-sale/cmd/app"
	"github.com/wesleyfebarretos/ticket-sale/config"
	"github.com/wesleyfebarretos/ticket-sale/infra/db"
	"github.com/wesleyfebarretos/ticket-sale/test/test_container"
)

var runningContainers = []*test_container.ContainerResult{}

func BeforeAll() {
	err := godotenv.Load("../../.env.test")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.Init()

	wg := &sync.WaitGroup{}
	runInParallel(wg, func() {
		runningContainers = append(runningContainers, test_container.SetupPG())
	})
	wg.Wait()

	db.Init()
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
