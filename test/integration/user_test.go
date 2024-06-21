package integration_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesleyfebarretos/ticket-sale/test/test_utils"
)

func TestGetAllUsers(t *testing.T) {
	t.Run("it should get all users and return 200", func(t *testing.T) {
		res, err := http.Get(fmt.Sprintf("%s/users", test_utils.Server.URL))
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
}
