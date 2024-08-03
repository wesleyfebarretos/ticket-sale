package test_data

import (
	"context"
	"testing"
	"time"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/creditcard_repository"
)

func NewCreditCard(t *testing.T, userID int32) creditcard_repository.CreateResponse {
	creditcard := creditcard_repository.New().Create(context.Background(), creditcard_repository.CreateParams{
		Name:             "Testing",
		Number:           "5574723384289379",
		Expiration:       time.Now().AddDate(3, 0, 0).UTC(),
		Priority:         1,
		NotifyExpiration: true,
		CreditcardTypeID: 1,
		CreditcardFlagID: 1,
		UserID:           userID,
	})

	return creditcard
}
