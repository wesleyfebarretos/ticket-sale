package test_utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http/httptest"
	"os"
	"reflect"
	"sync"
	"testing"

	"github.com/joho/godotenv"
	"github.com/wesleyfebarretos/ticket-sale/cmd/migrations/migration"
	"github.com/wesleyfebarretos/ticket-sale/config"
	"github.com/wesleyfebarretos/ticket-sale/infra/db"
	"github.com/wesleyfebarretos/ticket-sale/io/routes"
	"github.com/wesleyfebarretos/ticket-sale/middleware"
	"github.com/wesleyfebarretos/ticket-sale/repository/sqlc"
	"github.com/wesleyfebarretos/ticket-sale/test/test_container"
	"github.com/wesleyfebarretos/ticket-sale/utils"
)

var (
	runningContainers = []*test_container.ContainerResult{}
	UserTestPassword  = "123"
)

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
	migration.Up()

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

func CreateUser(role string) sqlc.GetUserWithPasswordByEmailRow {
	password, err := utils.HashPassword(UserTestPassword)
	if err != nil {
		log.Fatalf("could not hash password: %v", err)
	}

	newUser := sqlc.CreateUserParams{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@gmail.com",
		Password:  password,
		Role:      sqlc.Roles(role),
	}

	user, _ := db.Query.CreateUser(context.Background(), newUser)

	nUser, _ := db.Query.GetUserWithPasswordByEmail(context.Background(), user.Email)

	return nUser
}

func CreateUserAddress(userId int32) sqlc.UsersAddress {
	favorite := true
	complement := "Moon"
	postalCode := "Jupiter"
	addressType := "House"

	newAddress := sqlc.CreateUserAddressParams{
		Favorite:      &favorite,
		Complement:    &complement,
		PostalCode:    &postalCode,
		AddressType:   &addressType,
		StreetAddress: "Via Láctea",
		City:          "Dark Side",
		State:         "VL",
		Country:       "James Webb",
		UserID:        userId,
	}
	address, _ := db.Query.CreateUserAddress(context.Background(), newAddress)

	return address
}
