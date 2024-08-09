package user_service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/phone_types_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/user_address_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/user_phone_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/user_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/utils"
)

type CreateParams struct {
	User    user_repository.CreateParams
	Address user_address_repository.CreateParams
	Phone   user_phone_repository.CreateParams
}

type CreateResponse struct {
	User    user_repository.CreateResponse
	Address user_address_repository.CreateResponse
	Phone   user_phone_repository.CreateResponse
}

func GetAll(c *gin.Context) []user_repository.GetAllResponse {
	return user_repository.New().GetAll(c)
}

func GetOneById(c *gin.Context, id int32) *user_repository.GetOneByIdResponse {
	user := user_repository.New().GetOneById(c, user_repository.GetOneByIdParams{
		ID:   id,
		Role: roles_enum.USER,
	})
	if user == nil {
		panic(exception.NotFoundException(fmt.Sprintf("user of id %d not found", id)))
	}

	return user
}

func Create(c *gin.Context, body CreateParams) CreateResponse {
	return utils.WithTransaction(
		c,
		func(tx pgx.Tx) CreateResponse {
			userRepository := user_repository.New().WithTx(tx)

			user := userRepository.GetOneByEmail(c, body.User.Email)

			if user != nil {
				panic(exception.BadRequestException(fmt.Sprintf("email %s already registered", body.User.Email)))
			}

			hashPassword, err := utils.HashPassword(body.User.Password)
			if err != nil {
				panic(exception.InternalServerException(err.Error()))
			}

			body.User.Password = string(hashPassword)
			body.User.Role = roles_enum.USER

			createdUser := userRepository.Create(c, body.User)

			userAddressRepository := user_address_repository.New().WithTx(tx)

			body.Address.UserID = createdUser.ID

			newUserAddress := userAddressRepository.Create(c, body.Address)

			userPhoneRepository := user_phone_repository.New().WithTx(tx)

			body.Phone.Type = phone_types_enum.PHONE
			body.Phone.UserID = createdUser.ID

			newUserPhone := userPhoneRepository.Create(c, body.Phone)

			return CreateResponse{
				User:    createdUser,
				Address: newUserAddress,
				Phone:   newUserPhone,
			}
		},
	)
}

func Update(c *gin.Context, user user_repository.UpdateParams) {
	repository := user_repository.New()

	userExist := repository.CheckIfEmailExists(c, user_repository.CheckIfEmailExistsParams{
		Email: user.Email,
		ID:    user.ID,
	})

	if userExist != nil {
		panic(exception.BadRequestException(fmt.Sprintf("email %s already registered", user.Email)))
	}

	user.Role = roles_enum.USER

	repository.Update(c, user)
}

func GetFullProfile(c *gin.Context, id int32) *user_repository.GetProfileResponse {
	user := user_repository.New().GetProfile(c, id)
	if user == nil {
		panic(exception.NotFoundException(fmt.Sprintf("user of id %d not found", id)))
	}

	return user
}
