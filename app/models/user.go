package model

type User struct {
	ID    int
	Name  string
	Age   int
	Email string
}

var Users []User

func NewUser() User {
	return User{}
}

func (u User) Save() bool {
	Users = append(Users, u)
	return true
}
