package test_data

import (
	"context"
	"testing"
	"time"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/creditcard_repository"
)

func NewCreditCard(t *testing.T, userID int32) creditcard_repository.FinCreditcard {
	creditcard, err := repository.Creditcard.Create(context.Background(), creditcard_repository.CreateParams{
		Name:             "Testing",
		Number:           "5574723384289379",
		Expiration:       time.Now().AddDate(3, 0, 0).UTC(),
		Priority:         1,
		NotifyExpiration: true,
		CreditcardTypeID: 1,
		CreditcardFlagID: 1,
		UserID:           userID,
	})
	if err != nil {
		t.Errorf("error on create creditcard: %v", err)
	}

	return creditcard
}
