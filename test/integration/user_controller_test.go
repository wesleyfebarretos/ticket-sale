package integration_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesleyfebarretos/ticket-sale/internal/enum"
	user_controller "github.com/wesleyfebarretos/ticket-sale/io/http/controller/user"
	"github.com/wesleyfebarretos/ticket-sale/repository/sqlc"
)

func TestUsersController(t *testing.T) {
	expectedUser := &user_controller.CreateUserResponse{}

	t.Run("it should create a user", func(t *testing.T) {
		newUserRequest := user_controller.CreateUserRequest{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "johndoe@gmail.com",
			Password:  "123",
			Address: user_controller.AddressRequest{
				City:          "Orlando",
				State:         "FL",
				Favorite:      TPointer(true),
				Complement:    TPointer("Apartment 101"),
				PostalCode:    TPointer("32801"),
				AddressType:   TPointer("Residential"),
				StreetAddress: "123 Main St",
				Country:       "USA",
			},
		}

		res := TMakeRequest(t, http.MethodPost, "users", newUserRequest)

		newUserResponse := &user_controller.CreateUserResponse{}

		TDecode(t, res.Body, newUserResponse)

		expectedUser = &user_controller.CreateUserResponse{
			Id:        1,
			Role:      enum.USER_ROLE,
			FirstName: newUserResponse.FirstName,
			LastName:  newUserRequest.LastName,
			Email:     newUserRequest.Email,
			Address: &user_controller.AddressResponse{
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
				UpdatedAt:     newUserRequest.Address.UpdatedAt,
				CreatedAt:     newUserRequest.Address.CreatedAt,
			},
		}

		assert.Equal(t, http.StatusCreated, res.StatusCode)
		assert.Equal(t, expectedUser, newUserResponse)
	})

	t.Run("it should make sure that the user created has the role user", func(t *testing.T) {
		assert.Equal(t, expectedUser.Role, enum.USER_ROLE)
	})

	t.Run("it should get all", func(t *testing.T) {
		res := TMakeRequest(t, http.MethodGet, "users", nil)

		users := []sqlc.GetUsersRow{}

		TDecode(t, res.Body, &users)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, 1, len(users))
		assert.IsType(t, sqlc.GetUsersRow{}, users[0])
	})

	t.Run("it should get user full profile", func(t *testing.T) {
		res := TMakeRequest(t, http.MethodGet, "users/full-profile", nil)

		bSlice, err := io.ReadAll(res.Body)
		if err != nil {
			TErrorFatal(t, "could not read response body: %v", err)
		}

		defer res.Body.Close()

		userFullProfileRes := sqlc.GetUserFullProfileRow{}

		err = json.Unmarshal(bSlice, &userFullProfileRes)
		if err != nil {
			TErrorFatal(t, "could not parse to json: %v", err)
		}

		expectedJsonUser := map[string]interface{}{
			"id":        expectedUser.Id,
			"firstName": expectedUser.FirstName,
			"lastName":  expectedUser.LastName,
			"email":     expectedUser.Email,
			"role":      expectedUser.Role,
			"createdAt": userFullProfileRes.CreatedAt,
			"UpdatedAt": userFullProfileRes.UpdatedAt,
			"addresses": []map[string]interface{}{
				{
					"id":            expectedUser.Id,
					"userId":        expectedUser.Id,
					"streetAddress": expectedUser.Address.StreetAddress,
					"city":          expectedUser.Address.City,
					"complement":    expectedUser.Address.Complement,
					"state":         expectedUser.Address.State,
					"postalCode":    expectedUser.Address.PostalCode,
					"country":       expectedUser.Address.Country,
					"addressType":   expectedUser.Address.AddressType,
					"favorite":      expectedUser.Address.Favorite,
				},
			},
		}

		expectedBSlice, err := json.Marshal(expectedJsonUser)
		if err != nil {
			TErrorFatal(t, "could not marshal json to bytes: %v", err)
		}

		expectedUserFullProfile := sqlc.GetUserFullProfileRow{}

		err = json.Unmarshal(expectedBSlice, &expectedUserFullProfile)
		if err != nil {
			TErrorFatal(t, "could not parse to json: %v", err)
		}

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, expectedUserFullProfile, userFullProfileRes)
	})

	t.Run("it should get user by id", func(t *testing.T) {
		res := TMakeRequest(t, http.MethodGet, fmt.Sprintf("users/%d", expectedUser.Id), nil)

		userResponse := &sqlc.GetUserRow{}
		expectedUser := sqlc.GetUserRow{
			ID:    int32(expectedUser.Id),
			Email: expectedUser.Email,
			Role:  sqlc.Roles(expectedUser.Role),
		}

		TDecode(t, res.Body, &userResponse)

		assert.NotEmpty(t, userResponse)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, expectedUser.ID, userResponse.ID)
		assert.Equal(t, expectedUser.Email, userResponse.Email)
		assert.Equal(t, expectedUser.Role, userResponse.Role)
	})

	t.Run("it should be able to update an user", func(t *testing.T) {
		updateUser := user_controller.UpdateUserRequest{
			FirstName: "Update John",
			LastName:  "Update Doe",
			Email:     "updatejohndoe@gmail.com",
		}

		res := TMakeRequest(t, http.MethodPut, "users", updateUser)

		response := false

		TDecode(t, res.Body, &response)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.True(t, response)
	})

	t.Run("it should throw a user not found error", func(t *testing.T) {
		res := TMakeRequest(t, http.MethodGet, fmt.Sprintf("users/%d", 100), nil)

		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
}
