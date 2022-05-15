package user

type Storage interface {
	CreateUser(user *User) (*User, error)
	GetUserById(userId int) (*User, error)
	GetUserByEmail(userEmail string) (*User, error)
}
