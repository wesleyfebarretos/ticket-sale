package integration_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/handler/user_handler"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/middleware"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/tests/test_data"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/tests/test_utils"
)

func TestUsersHandler(t *testing.T) {
	t.Run("it should create a user", TRun(func(t *testing.T) {
		newUserRequest := user_handler.CreateRequestDto{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "johndoe@gmail.com",
			Password:  "123",
			Address: user_handler.AddressRequestDto{
				City:          "Orlando",
				State:         "FL",
				Favorite:      TPointer(true),
				Complement:    TPointer("Apartment 101"),
				PostalCode:    TPointer("32801"),
				AddressType:   TPointer("Residential"),
				StreetAddress: "123 Main St",
				Country:       "USA",
			},
			Phone: user_handler.PhoneRequestDto{
				Ddd:    "407",
				Number: "5551234",
			},
		}
		res := TMakeRequest(t, http.MethodPost, "users", newUserRequest)

		newUserResponse := &user_handler.CreateResponseDto{}

		test_utils.Decode(t, res.Body, &newUserResponse)

		expectedUser := &user_handler.CreateResponseDto{
			Id:        newUserResponse.Id,
			Role:      roles_enum.USER,
			FirstName: newUserResponse.FirstName,
			LastName:  newUserRequest.LastName,
			Email:     newUserRequest.Email,
			Address: user_handler.AddressResponseDto{
				ID:            newUserResponse.Address.ID,
				UserID:        newUserResponse.Address.UserID,
				Country:       newUserRequest.Address.Country,
				StreetAddress: newUserRequest.Address.StreetAddress,
				AddressType:   newUserRequest.Address.AddressType,
				PostalCode:    newUserRequest.Address.PostalCode,
				Complement:    newUserRequest.Address.Complement,
				Favorite:      newUserRequest.Address.Favorite,
				State:         newUserRequest.Address.State,
				City:          newUserRequest.Address.City,
				UpdatedAt:     newUserResponse.Address.UpdatedAt,
				CreatedAt:     newUserResponse.Address.CreatedAt,
			},
			Phone: user_handler.PhoneResponseDto{
				ID:     newUserResponse.Phone.ID,
				UserID: int32(newUserResponse.Id),
				Ddd:    newUserResponse.Phone.Ddd,
				Number: newUserResponse.Phone.Number,
			},
		}

		assert.Equal(t, http.StatusCreated, res.StatusCode)
		assert.Equal(t, expectedUser, newUserResponse)
		assert.Equal(t, newUserResponse.Role, roles_enum.USER)
	}))

	t.Run("it should login", TRun(func(t *testing.T) {
		user := test_data.NewUser(roles_enum.USER)

		loginRequest := middleware.SignInRequest{
			Email:    user.Email,
			Password: test_data.UserTestPassword,
		}

		res := TMakeRequest(t, http.MethodPost, "auth", loginRequest)

		responseBody := middleware.SignInResponse{}

		test_utils.Decode(t, res.Body, &responseBody)

		jwtTimeOut := middleware.BuildJwtTimeOut()

		jwtTimeOutMinusOne := time.Now().Add(jwtTimeOut - (time.Duration(1) * time.Minute))

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.NotEmpty(t, responseBody.Token)
		assert.IsType(t, "", responseBody.Token)
		assert.True(t, responseBody.Expire.After(jwtTimeOutMinusOne), "Expected expiration time to be after actual expiration time")
	}))

	t.Run("it should get all", TRun(func(t *testing.T) {
		user := test_data.NewUser(roles_enum.USER)
		TSetCookieWithUser(t, user)

		res := TMakeRequest(t, http.MethodGet, "users", nil)

		users := []user_handler.GetAllResponseDto{}

		test_utils.Decode(t, res.Body, &users)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, 1, len(users))
	}))

	t.Run("it should get user full profile", TRun(func(t *testing.T) {
		user := test_data.NewUser(roles_enum.USER)
		TSetCookieWithUser(t, user)
		userAddress := test_data.NewUserAddress(user.ID)
		userPhone := test_data.NewUserPhone(user.ID)

		res := TMakeRequest(t, http.MethodGet, "users/full-profile", nil)

		bSlice, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read response body: %v", err)
		}

		defer res.Body.Close()

		userFullProfileRes := user_handler.GetProfileResponseDto{}

		err = json.Unmarshal(bSlice, &userFullProfileRes)
		if err != nil {
			t.Fatalf("could not parse to json: %v", err)
		}

		expectedJsonUser := map[string]interface{}{
			"id":        user.ID,
			"firstName": user.FirstName,
			"lastName":  user.LastName,
			"email":     user.Email,
			"role":      user.Role,
			"createdAt": user.CreatedAt,
			"UpdatedAt": user.UpdatedAt,
			"addresses": []map[string]interface{}{
				{
					"id":            userAddress.ID,
					"userId":        userAddress.UserID,
					"streetAddress": userAddress.StreetAddress,
					"city":          userAddress.City,
					"complement":    userAddress.Complement,
					"state":         userAddress.State,
					"postalCode":    userAddress.PostalCode,
					"country":       userAddress.Country,
					"addressType":   userAddress.AddressType,
					"favorite":      userAddress.Favorite,
				},
			},
			"phones": []map[string]interface{}{
				{
					"id":     userPhone.ID,
					"userId": userPhone.UserID,
					"ddd":    userPhone.Ddd,
					"number": userPhone.Number,
					"type":   userPhone.Type,
				},
			},
		}

		expectedBSlice, err := json.Marshal(expectedJsonUser)
		if err != nil {
			t.Fatalf("could not marshal json to bytes: %v", err)
		}

		expectedUserFullProfile := user_handler.GetProfileResponseDto{}

		err = json.Unmarshal(expectedBSlice, &expectedUserFullProfile)
		if err != nil {
			t.Fatalf("could not parse to json: %v", err)
		}

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, expectedUserFullProfile, userFullProfileRes)
	}))

	t.Run("it should get user by id", TRun(func(t *testing.T) {
		user := test_data.NewUser(roles_enum.USER)
		TSetCookieWithUser(t, user)

		res := TMakeRequest(t, http.MethodGet, fmt.Sprintf("users/%d", user.ID), nil)

		userResponse := &user_handler.GetOneByIdResponseDto{}
		expectedUser := user_handler.GetOneByIdResponseDto{
			Id:    user.ID,
			Email: user.Email,
			Role:  string(user.Role),
		}

		test_utils.Decode(t, res.Body, &userResponse)

		assert.NotEmpty(t, userResponse)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, expectedUser.Id, userResponse.Id)
		assert.Equal(t, expectedUser.Email, userResponse.Email)
		assert.Equal(t, expectedUser.Role, userResponse.Role)
	}))

	t.Run("it should be able to update an user", TRun(func(t *testing.T) {
		user := test_data.NewUser(roles_enum.USER)
		TSetCookieWithUser(t, user)

		updateUser := user_handler.UpdateRequestDto{
			FirstName: "Update John",
			LastName:  "Update Doe",
			Email:     "updatejohndoe@gmail.com",
		}

		res := TMakeRequest(t, http.MethodPut, "users", updateUser)

		response := false

		test_utils.Decode(t, res.Body, &response)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.True(t, response)
	}))

	t.Run("it should throw a user not found error", TRun(func(t *testing.T) {
		user := test_data.NewUser(roles_enum.USER)
		TSetCookieWithUser(t, user)
		res := TMakeRequest(t, http.MethodGet, fmt.Sprintf("users/%d", 100), nil)

		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	}))
}
