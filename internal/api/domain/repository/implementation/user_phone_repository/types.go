package user_phone_repository

type CreateParams struct {
	UserID int32  `json:"userId"`
	Ddd    string `json:"ddd"`
	Number string `json:"number"`
	Type   string `json:"type"`
}

type CreateResponse struct {
	ID     int32  `json:"id"`
	UserID int32  `json:"userId"`
	Ddd    string `json:"ddd"`
	Number string `json:"number"`
	Type   string `json:"type"`
}
