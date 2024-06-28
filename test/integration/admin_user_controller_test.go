package integration_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesleyfebarretos/ticket-sale/internal/enum"
	"github.com/wesleyfebarretos/ticket-sale/io/http/controller/admin_user_controller"
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
			Password:  "123",
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
}
