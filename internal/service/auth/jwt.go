package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wesleyfebarretos/ticket-sale/config"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/types"
)

func CreateJWT(secret []byte, claims types.JWTPayload) string {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        claims.ID,
		"role":      claims.Role,
		"expiredAt": expiration,
	})

	strToken, err := token.SignedString(secret)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return strToken
}
