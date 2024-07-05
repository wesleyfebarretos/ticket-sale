package db_util

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/infra/db"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
)

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
