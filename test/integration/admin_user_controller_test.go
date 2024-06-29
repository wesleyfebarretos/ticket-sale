package integration_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"github.com/wesleyfebarretos/ticket-sale/internal/enum"
	"github.com/wesleyfebarretos/ticket-sale/io/http/controller/admin_user_controller"
	"github.com/wesleyfebarretos/ticket-sale/repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_user_repository"
	"github.com/wesleyfebarretos/ticket-sale/test/test_utils"
)

func TestAdminUsersController(t *testing.T) {
	t.Run("it should create an admin user", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(enum.ADMIN_ROLE)
		TSetCookieWithUser(t, adminUser)

		newAdminUser := admin_user_controller.CreateRequestDto{
			Email:     "adminjohndoe@gmail.com",
			FirstName: "John",
			LastName:  "Doe",
			Password:  "123456",
		}

		res := TMakeRequest(t, http.MethodPost, "admin/users", newAdminUser)

		newAdminUserResponse := admin_user_controller.CreateResponseDto{}

		test_utils.Decode(t, res.Body, &newAdminUserResponse)

		expectedAdminUser := admin_user_controller.CreateResponseDto{
			ID:        newAdminUserResponse.ID,
			FirstName: newAdminUser.FirstName,
			LastName:  newAdminUser.LastName,
			Email:     newAdminUser.Email,
			Role:      enum.ADMIN_ROLE,
			CreatedAt: newAdminUserResponse.CreatedAt,
			UpdatedAt: newAdminUserResponse.UpdatedAt,
		}

		assert.Equal(t, http.StatusCreated, res.StatusCode)
		assert.Equal(t, expectedAdminUser, newAdminUserResponse)
		assert.Equal(t, enum.ADMIN_ROLE, newAdminUserResponse.Role)
		assert.GreaterOrEqual(t, newAdminUserResponse.ID, int32(1))
	}))

	t.Run("it should update an user", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(enum.ADMIN_ROLE)
		TSetCookieWithUser(t, adminUser)

		updateAdminUser := admin_user_controller.UpdateRequestDto{
			FirstName: "updateJohn",
			LastName:  "updateDoe",
			Email:     "updateadminjohndoe@gmail.com",
			Role:      enum.ADMIN_ROLE,
		}

		res := TMakeRequest(t, http.MethodPut, fmt.Sprintf("admin/users/%d", adminUser.ID), updateAdminUser)

		updateAdminUserResponse := false

		test_utils.Decode(t, res.Body, &updateAdminUserResponse)

		updatedUser, err := repository.AdminUser.GetOneById(context.Background(), admin_user_repository.GetOneByIdParams{
			ID:   adminUser.ID,
			Role: enum.ADMIN_ROLE,
		})
		if err != nil {
			t.Errorf("updated admin user of id %d not found", adminUser.ID)
		}

		updatedAdminUserNewData := admin_user_controller.UpdateRequestDto{
			FirstName: updatedUser.FirstName,
			LastName:  updatedUser.LastName,
			Email:     updatedUser.Email,
			Role:      string(updatedUser.Role),
		}

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, true, updateAdminUserResponse)
		assert.Equal(t, updateAdminUser, updatedAdminUserNewData)
	}))

	t.Run("it should delete an user", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(enum.ADMIN_ROLE)
		TSetCookieWithUser(t, adminUser)
		ctx := context.Background()

		beforeDeleteAdminUser, err := repository.AdminUser.GetOneById(ctx, admin_user_repository.GetOneByIdParams{
			ID:   adminUser.ID,
			Role: enum.ADMIN_ROLE,
		})
		if err != nil && err != pgx.ErrNoRows {
			t.Error(err)
		}

		res := TMakeRequest(t, http.MethodDelete, fmt.Sprintf("admin/users/%d", adminUser.ID), nil)

		deleteAdminUsersResponse := false

		test_utils.Decode(t, res.Body, &deleteAdminUsersResponse)

		_, err = repository.AdminUser.GetOneById(ctx, admin_user_repository.GetOneByIdParams{
			ID:   adminUser.ID,
			Role: enum.ADMIN_ROLE,
		})

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, true, deleteAdminUsersResponse)
		assert.GreaterOrEqual(t, beforeDeleteAdminUser.ID, int32(1))
		assert.Error(t, err)
		assert.ErrorContains(t, err, "no rows in result set")
	}))

	t.Run("it should get one by id", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(enum.ADMIN_ROLE)
		TSetCookieWithUser(t, adminUser)

		res := TMakeRequest(t, http.MethodGet, fmt.Sprintf("admin/users/%d", adminUser.ID), nil)

		adminUserResponse := admin_user_controller.GetOneByIdResponseDto{}

		test_utils.Decode(t, res.Body, &adminUserResponse)

		expectedAdminUser := admin_user_controller.GetOneByIdResponseDto{
			ID:        adminUser.ID,
			FirstName: adminUser.FirstName,
			LastName:  adminUser.LastName,
			Email:     adminUser.Email,
			Role:      enum.ADMIN_ROLE,
			CreatedAt: adminUser.CreatedAt,
			UpdatedAt: adminUser.UpdatedAt,
		}

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, expectedAdminUser, adminUserResponse)
	}))

	t.Run("it should get one by email", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(enum.ADMIN_ROLE)
		TSetCookieWithUser(t, adminUser)

		res := TMakeRequest(t, http.MethodPost, "admin/users/get-by-email", admin_user_controller.GetOneByEmailRequestDto{
			Email: adminUser.Email,
		})

		adminUserResponse := admin_user_controller.GetOneByEmailResponseDto{}

		test_utils.Decode(t, res.Body, &adminUserResponse)

		expectedAdminUser := admin_user_controller.GetOneByEmailResponseDto{
			ID:        adminUser.ID,
			FirstName: adminUser.FirstName,
			LastName:  adminUser.LastName,
			Email:     adminUser.Email,
			Role:      enum.ADMIN_ROLE,
			CreatedAt: adminUser.CreatedAt,
			UpdatedAt: adminUser.UpdatedAt,
		}

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, expectedAdminUser, adminUserResponse)
	}))

	t.Run("it should get all", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(enum.ADMIN_ROLE)
		TSetCookieWithUser(t, adminUser)

		ctx := context.Background()
		for i := 0; i < 10; i++ {
			repository.AdminUser.Create(ctx, admin_user_repository.CreateParams{
				FirstName: "John",
				LastName:  "Doe",
				Email:     fmt.Sprintf("johndoefor%d@gmail.com", i),
				Password:  "123",
				Role:      enum.ADMIN_ROLE,
			})
		}

		res := TMakeRequest(t, http.MethodGet, "admin/users", nil)

		adminUsersResponse := []admin_user_controller.GetAllResponseDto{}

		test_utils.Decode(t, res.Body, &adminUsersResponse)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.GreaterOrEqual(t, len(adminUsersResponse), 10)
		assert.LessOrEqual(t, len(adminUsersResponse), 12)
	}))

	t.Run("it should fail with an unauthorized error", TRun(func(t *testing.T) {
		user := test_utils.CreateUser(enum.USER_ROLE)

		TSetCookieWithUser(t, user)

		res := TMakeRequest(t, http.MethodGet, "admin/users", nil)

		assert.Equal(t, http.StatusForbidden, res.StatusCode)
	}))

	t.Run("it should try to create and fail with an first name required error", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(enum.ADMIN_ROLE)

		TSetCookieWithUser(t, adminUser)

		newAdminUser := admin_user_controller.CreateRequestDto{
			FirstName: "",
			LastName:  "Doe",
			Email:     "johndoe@gmail.com",
			Password:  "123456",
		}

		res := TMakeRequest(t, http.MethodPost, "admin/users", newAdminUser)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	}))

	t.Run("it should try to create and fail with an last name required error", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(enum.ADMIN_ROLE)

		TSetCookieWithUser(t, adminUser)

		newAdminUser := admin_user_controller.CreateRequestDto{
			FirstName: "John",
			LastName:  "",
			Email:     "johndoe@gmail.com",
			Password:  "123456",
		}

		res := TMakeRequest(t, http.MethodPost, "admin/users", newAdminUser)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	}))

	t.Run("it should try to create and fail with an email required error", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(enum.ADMIN_ROLE)

		TSetCookieWithUser(t, adminUser)

		newAdminUser := admin_user_controller.CreateRequestDto{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "",
			Password:  "123456",
		}

		res := TMakeRequest(t, http.MethodPost, "admin/users", newAdminUser)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	}))

	t.Run("it should try to create and fail with an password required error", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(enum.ADMIN_ROLE)

		TSetCookieWithUser(t, adminUser)

		newAdminUser := admin_user_controller.CreateRequestDto{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "johndoe@gmail.com",
			Password:  "",
		}

		res := TMakeRequest(t, http.MethodPost, "admin/users", newAdminUser)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	}))

	t.Run("it should try to create and fail with an first name min length error", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(enum.ADMIN_ROLE)

		TSetCookieWithUser(t, adminUser)

		newAdminUser := admin_user_controller.CreateRequestDto{
			FirstName: "Jo",
			LastName:  "Doe",
			Email:     "johndoe@gmail.com",
			Password:  "123456",
		}

		res := TMakeRequest(t, http.MethodPost, "admin/users", newAdminUser)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	}))

	t.Run("it should try to create and fail with an password min length error", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(enum.ADMIN_ROLE)

		TSetCookieWithUser(t, adminUser)

		newAdminUser := admin_user_controller.CreateRequestDto{
			FirstName: "John",
			LastName:  "D",
			Email:     "johndoe@gmail.com",
			Password:  "12345",
		}

		res := TMakeRequest(t, http.MethodPost, "admin/users", newAdminUser)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	}))

	t.Run("it should try to create and fail with an first name max length error", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(enum.ADMIN_ROLE)

		TSetCookieWithUser(t, adminUser)

		firstName := ""

		for i := 0; i < 51; i++ {
			firstName += "a"
		}

		newAdminUser := admin_user_controller.CreateRequestDto{
			FirstName: firstName,
			LastName:  "Doe",
			Email:     "johndoe@gmail.com",
			Password:  "123456",
		}

		res := TMakeRequest(t, http.MethodPost, "admin/users", newAdminUser)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	}))

	t.Run("it should try to create and fail with an last name max length error", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(enum.ADMIN_ROLE)

		TSetCookieWithUser(t, adminUser)

		lastName := ""

		for i := 0; i < 51; i++ {
			lastName += "a"
		}

		newAdminUser := admin_user_controller.CreateRequestDto{
			FirstName: "John",
			LastName:  lastName,
			Email:     "johndoe@gmail.com",
			Password:  "123456",
		}

		res := TMakeRequest(t, http.MethodPost, "admin/users", newAdminUser)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	}))

	t.Run("it should try to update and fail with an first name required error", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(enum.ADMIN_ROLE)

		TSetCookieWithUser(t, adminUser)

		updateAdminUser := admin_user_controller.UpdateRequestDto{
			FirstName: "",
			LastName:  "Doe",
			Email:     "johndoe@gmail.com",
			Role:      enum.ADMIN_ROLE,
		}

		res := TMakeRequest(t, http.MethodPut, fmt.Sprintf("admin/users/%d", adminUser.ID), updateAdminUser)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	}))

	t.Run("it should try to update and fail with an last name required error", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(enum.ADMIN_ROLE)

		TSetCookieWithUser(t, adminUser)

		updateAdminUser := admin_user_controller.UpdateRequestDto{
			FirstName: "John",
			LastName:  "",
			Email:     "johndoe@gmail.com",
			Role:      enum.ADMIN_ROLE,
		}

		res := TMakeRequest(t, http.MethodPut, fmt.Sprintf("admin/users/%d", adminUser.ID), updateAdminUser)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	}))

	t.Run("it should try to update and fail with an email required error", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(enum.ADMIN_ROLE)

		TSetCookieWithUser(t, adminUser)

		updateAdminUser := admin_user_controller.UpdateRequestDto{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "",
			Role:      enum.ADMIN_ROLE,
		}

		res := TMakeRequest(t, http.MethodPut, fmt.Sprintf("admin/users/%d", adminUser.ID), updateAdminUser)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	}))

	t.Run("it should try to update and fail with an role required error", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(enum.ADMIN_ROLE)

		TSetCookieWithUser(t, adminUser)

		updateAdminUser := admin_user_controller.UpdateRequestDto{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "johndoe@gmail.com",
		}

		res := TMakeRequest(t, http.MethodPut, fmt.Sprintf("admin/users/%d", adminUser.ID), updateAdminUser)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	}))

	t.Run("it should try to update and fail with an first name min length error", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(enum.ADMIN_ROLE)

		TSetCookieWithUser(t, adminUser)

		updateAdminUser := admin_user_controller.UpdateRequestDto{
			FirstName: "Jo",
			LastName:  "Doe",
			Email:     "johndoe@gmail.com",
			Role:      enum.ADMIN_ROLE,
		}

		res := TMakeRequest(t, http.MethodPut, fmt.Sprintf("admin/users/%d", adminUser.ID), updateAdminUser)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	}))

	t.Run("it should try to update and fail with an last name min length error", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(enum.ADMIN_ROLE)

		TSetCookieWithUser(t, adminUser)

		updateAdminUser := admin_user_controller.UpdateRequestDto{
			FirstName: "John",
			LastName:  "D",
			Email:     "johndoe@gmail.com",
			Role:      enum.ADMIN_ROLE,
		}

		res := TMakeRequest(t, http.MethodPut, fmt.Sprintf("admin/users/%d", adminUser.ID), updateAdminUser)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	}))

	t.Run("it should try to update and fail with an invid role error", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(enum.ADMIN_ROLE)

		TSetCookieWithUser(t, adminUser)

		updateAdminUser := admin_user_controller.UpdateRequestDto{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "johndoe@gmail.com",
			Role:      "teste",
		}

		res := TMakeRequest(t, http.MethodPut, fmt.Sprintf("admin/users/%d", adminUser.ID), updateAdminUser)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	}))

	t.Run("it should try to update and fail with an first name max length error", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(enum.ADMIN_ROLE)

		TSetCookieWithUser(t, adminUser)

		firstName := ""

		for i := 0; i < 51; i++ {
			firstName += "a"
		}

		updateAdminUser := admin_user_controller.UpdateRequestDto{
			FirstName: firstName,
			LastName:  "Doe",
			Email:     "johndoe@gmail.com",
			Role:      enum.ADMIN_ROLE,
		}

		res := TMakeRequest(t, http.MethodPut, fmt.Sprintf("admin/users/%d", adminUser.ID), updateAdminUser)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	}))

	t.Run("it should try to update and fail with an last name max length error", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(enum.ADMIN_ROLE)

		TSetCookieWithUser(t, adminUser)

		lastName := ""

		for i := 0; i < 51; i++ {
			lastName += "a"
		}

		updateAdminUser := admin_user_controller.UpdateRequestDto{
			FirstName: "John",
			LastName:  lastName,
			Email:     "johndoe@gmail.com",
			Role:      enum.ADMIN_ROLE,
		}

		res := TMakeRequest(t, http.MethodPut, fmt.Sprintf("admin/users/%d", adminUser.ID), updateAdminUser)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	}))

	t.Run("it should fail with a not found admin user error", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(enum.ADMIN_ROLE)
		TSetCookieWithUser(t, adminUser)

		res := TMakeRequest(t, http.MethodGet, fmt.Sprintf("admin/users/%d", 100), nil)

		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	}))
}
