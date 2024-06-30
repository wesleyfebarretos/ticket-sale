package admin_auth_controller

// REQUESTS
type SignInRequestDto struct {
	Email    string `json:"email" binding:"required,email" example:"ticketsale@gmail.com"`
	Password string `json:"password" binding:"required" example:"123"`
}

// RESPONSES
type SignInResponseDto struct {
	Expire string `json:"expire" example:"2024-06-30T20:46:13-03:00"`
	Token  string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTk3OTExNzMsImlkIjozLCJvcmlnX2lhdCI6MTcxOTcwNDc3Mywicm9sZSI6InVzZXIifQ.c8HuyRAxgNDC4FavwQ_mv-qWOm4Ch6--1-kSQEmK4x0"`
	Code   int    `json:"code" example:"200"`
}
