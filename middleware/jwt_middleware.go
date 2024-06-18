package middleware

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/config"
	"github.com/wesleyfebarretos/ticket-sale/infra/db"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/repository/sqlc"
	"github.com/wesleyfebarretos/ticket-sale/utils"
)

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type SignInResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type UserClaims struct {
	Role string
	Id   int32
}

type AuthenticationError struct {
	Message string
	Code    int
}

var Jwt = jwtMiddleware()

var jwtTimeout = time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)

const IDENTITY_KEY = "user"

func jwtMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           config.Envs.PublicHost,
		Key:             []byte(config.Envs.JWTSecret),
		Timeout:         jwtTimeout,
		Authenticator:   loginHandler,
		PayloadFunc:     payloadHandler,
		IdentityHandler: identityHandler,
		Authorizator:    autorizatorHandler,
		Unauthorized:    unauthorizedHandler,
		IdentityKey:     IDENTITY_KEY,
		TimeFunc:        time.Now,
		SendCookie:      true,
		SecureCookie:    false, // non HTTPS dev environments
		CookieHTTPOnly:  true,  // JS can't modify
		CookieDomain:    config.Envs.PublicHost,
		CookieName:      config.Envs.CookieName, // default jwt
		TokenLookup:     fmt.Sprintf("cookie:%s", config.Envs.CookieName),
		CookieSameSite:  http.SameSiteDefaultMode, // SameSiteDefaultMode, SameSiteLaxMode, SameSiteStrictMode, SameSiteNoneMode
	})
	if err != nil {
		log.Fatal("JWT Initialization Error: " + err.Error())
	}

	return authMiddleware
}

func loginHandler(c *gin.Context) (interface{}, error) {
	body := SignInRequest{}
	readBody(c, &body)
	user, err := db.Query.GetUserWithPasswordByEmail(c, body.Email)
	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	if !utils.ComparePassword(user.Password, body.Password) {
		return nil, jwt.ErrFailedAuthentication
	}

	return user, nil
}

func readBody(c *gin.Context, body any) {
	err := c.ShouldBindJSON(&body)
	if err == io.EOF {
		panic(exception.BadRequestException("empty request body"))
	}
	if err != nil {
		panic(exception.BadRequestException(err.Error()))
	}
}

func payloadHandler(data interface{}) jwt.MapClaims {
	user := data.(sqlc.GetUserWithPasswordByEmailRow)

	return jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
	}
}

func identityHandler(c *gin.Context) interface{} {
	return ExtractClaims(c)
}

func autorizatorHandler(data interface{}, c *gin.Context) bool {
	return true
}

func unauthorizedHandler(c *gin.Context, code int, message string) {
	c.JSON(code, AuthenticationError{
		Code:    code,
		Message: message,
	})
}

func ExtractClaims(c *gin.Context) *UserClaims {
	claims := jwt.ExtractClaims(c)
	return &UserClaims{
		Id:   int32(claims["id"].(float64)),
		Role: claims["role"].(string),
	}
}
