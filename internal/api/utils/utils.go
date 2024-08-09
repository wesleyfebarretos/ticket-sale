package utils

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/infra/db"
)

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func WithTransaction[R any](c context.Context, fn func(pgx.Tx) R) R {
	tx, err := db.Conn.Begin(c)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}
	defer tx.Rollback(c)

	response := fn(tx)

	if err := tx.Commit(c); err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return response
}

func MaskCreditcardNumber(number string) string {
	firstFourDigits := number[:4]
	lastFourDigits := number[len(number)-4:]
	masked := strings.Repeat("*", len(number)-8)

	return firstFourDigits + masked + lastFourDigits
}
