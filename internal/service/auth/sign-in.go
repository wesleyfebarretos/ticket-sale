package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/config"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/types"
	"github.com/wesleyfebarretos/ticket-sale/io/dto"
	"github.com/wesleyfebarretos/ticket-sale/repository/sqlc"
)

func UserSignIn(c *gin.Context, query *sqlc.Queries, body dto.SignInRequest) *dto.SignInResponse {
	user, err := query.GetUserWithPasswordByEmail(c, body.Email)
	if err == pgx.ErrNoRows {
		panic(exception.BadRequestException("email or password invalid"))
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	if !ComparePassword(user.Password, body.Password) {
		panic(exception.BadRequestException("email or password invalid"))
	}

	token := CreateJWT([]byte(config.Envs.JWTSecret), types.JWTPayload{
		ID:   int(user.ID),
		Role: string(user.Role),
	})

	response := &dto.SignInResponse{
		Email:    body.Email,
		Password: body.Password,
		Token:    token,
	}

	return response
}
