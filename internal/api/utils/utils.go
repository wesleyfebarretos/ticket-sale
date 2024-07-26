package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/infra/db"
	"golang.org/x/crypto/bcrypt"
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

func WithTransaction[R any](c *gin.Context, fn func(pgx.Tx) R) R {
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
