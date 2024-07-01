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
	"github.com/wesleyfebarretos/ticket-sale/internal/enum/phone_types_enum"
	"github.com/wesleyfebarretos/ticket-sale/io/routes"
	"github.com/wesleyfebarretos/ticket-sale/middleware"
	"github.com/wesleyfebarretos/ticket-sale/repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/users_addresses_repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/users_phones_repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/users_repository"
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
	migration.UpTables()
	repository.Bind()

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

func CreateUser(role string) users_repository.GetOneWithPasswordByEmailRow {
	password, err := utils.HashPassword(UserTestPassword)
	if err != nil {
		log.Fatalf("could not hash password: %v", err)
	}

	newUser := users_repository.CreateParams{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoetest@gmail.com",
		Password:  password,
		Role:      users_repository.Roles(role),
	}

	user, _ := repository.Users.Create(context.Background(), newUser)

	nUser, _ := repository.Users.GetOneWithPasswordByEmail(context.Background(), user.Email)

	return nUser
}

func CreateUserAddress(userId int32) users_addresses_repository.UsersAddress {
	favorite := true
	complement := "Moon"
	postalCode := "Jupiter"
	addressType := "House"

	newAddress := users_addresses_repository.CreateParams{
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

	address, _ := repository.UsersAdresses.Create(context.Background(), newAddress)

	return address
}

func CreateUserPhone(userID int32) users_phones_repository.UsersPhone {
	newPhone := users_phones_repository.CreateParams{
		UserID: userID,
		Ddd:    "021",
		Number: "999999999",
		Type:   phone_types_enum.PHONE,
	}

	userPhone, _ := repository.UsersPhones.Create(context.Background(), newPhone)

	return userPhone
}

func DebugResponse(body io.Reader) {
	b, _ := io.ReadAll(body)
	fmt.Println(string(b))
	os.Exit(1)
}
