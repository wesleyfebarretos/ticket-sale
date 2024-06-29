package admin_user_controller

import "time"

// REQUESTS
type GetOneByEmailRequestDto struct {
	Email string `json:"email" binding:"required,email"`
}

type CreateRequestDto struct {
	FirstName string `json:"firstName" binding:"required,max=50,min=3"`
	LastName  string `json:"lastName" binding:"required,max=50,min=2"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,max=1000,min=6"`
}

type UpdateRequestDto struct {
	FirstName string `json:"firstName" binding:"required,max=50,min=3"`
	LastName  string `json:"lastName" binding:"required,max=50,min=2"`
	Email     string `json:"email" binding:"required,email"`
	Role      string `json:"role" binding:"required,oneof=admin user webservice'"`
}

// RESPONSES
type GetOneByIdResponseDto struct {
	ID        int32      `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	Role      string     `json:"role"`
	UpdatedAt *time.Time `json:"updatedAt"`
	CreatedAt time.Time  `json:"createdAt"`
}

type GetOneByEmailResponseDto struct {
	ID        int32      `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type CreateResponseDto struct {
	ID        int32      `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type GetAllResponseDto struct {
	ID        int32      `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
