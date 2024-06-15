package entities

type UserEntity struct {
	FirstName string
	LastName  string
	email     string
	password  string
	Role      int
	Id        int
}

func (u *UserEntity) TableName() string {
	return "user"
}
