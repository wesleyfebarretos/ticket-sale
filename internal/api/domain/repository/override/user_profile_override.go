package override

type UserProfilePhone struct {
	ID     int32  `json:"id"`
	UserID int32  `json:"userId"`
	DDD    string `json:"ddd"`
	Number string `json:"number"`
	Type   string `json:"type"`
}

type UserProfileAddress struct {
	ID            int32  `json:"id"`
	UserID        int32  `json:"userId"`
	StreetAddress string `json:"streetAddress"`
	City          string `json:"city"`
	Complement    string `json:"complement"`
	State         string `json:"state"`
	PostalCode    string `json:"postalCode"`
	Country       string `json:"country"`
	AddressType   string `json:"addressType"`
	Favorite      bool   `json:"favorite"`
}
