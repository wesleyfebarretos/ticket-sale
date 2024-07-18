package admin_user_controller

import "time"

// REQUESTS
type GetOneByEmailRequestDto struct {
	Email string `json:"email" binding:"required,email" example:"adminjohndoe@gmail.com"`
}

type CreateRequestDto struct {
	FirstName string `json:"firstName" binding:"required,max=50,min=3" example:"Admin John"`
	LastName  string `json:"lastName" binding:"required,max=50,min=2" example:"Doe"`
	Email     string `json:"email" binding:"required,email" example:"adminjohndoe@gmail.com"`
	Password  string `json:"password" binding:"required,max=1000,min=6" example:"123456"`
}

type UpdateRequestDto struct {
	FirstName string `json:"firstName" binding:"required,max=50,min=3" example:"Admin John"`
	LastName  string `json:"lastName" binding:"required,max=50,min=2" example:"Doe"`
	Email     string `json:"email" binding:"required,email" example:"adminjohndoe@gmail.com"`
	Role      string `json:"role" binding:"required,oneof=admin user webservice'" enums:"admin,user,webservice"`
}

// RESPONSES
type GetOneByIdResponseDto struct {
	ID        int32     `json:"id" example:"1"`
	FirstName string    `json:"firstName" example:"Admin John"`
	LastName  string    `json:"lastName" example:"Doe"`
	Email     string    `json:"email" example:"adminjohndoe@gmail.com"`
	Role      string    `json:"role" example:"admin"`
	CreatedAt time.Time `json:"createdAt" example:"2024-06-29T06:29:44.999929Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2024-06-29T06:29:44.999929Z"`
}

type GetOneByEmailResponseDto struct {
	ID        int32     `json:"id" example:"1"`
	FirstName string    `json:"firstName" example:"Admin John"`
	LastName  string    `json:"lastName" example:"Doe"`
	Email     string    `json:"email" example:"adminjohndoe@gmail.com"`
	Role      string    `json:"role" example:"admin"`
	CreatedAt time.Time `json:"createdAt" example:"2024-06-29T06:29:44.999929Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2024-06-29T06:29:44.999929Z"`
}

type CreateResponseDto struct {
	ID        int32     `json:"id" example:"2"`
	FirstName string    `json:"firstName" binding:"required,max=50,min=3" example:"Admin John"`
	LastName  string    `json:"lastName" binding:"required,max=50,min=2" example:"Doe"`
	Email     string    `json:"email" binding:"required,email" example:"adminjohndoe@gmail.com"`
	Role      string    `json:"role" example:"admin"`
	CreatedAt time.Time `json:"createdAt" example:"2024-06-29T06:29:44.999929Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2024-06-29T06:29:44.999929Z"`
}

type GetAllResponseDto struct {
	ID        int32     `json:"id" example:"1"`
	FirstName string    `json:"firstName" example:"John"`
	LastName  string    `json:"lastName" example:"Doe"`
	Email     string    `json:"email" example:"johndoe@gmail.com"`
	Role      string    `json:"role" example:"admin"`
	CreatedAt time.Time `json:"createdAt" example:"2024-06-29T06:29:44.999929Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2024-06-29T06:29:44.999929Z"`
}
