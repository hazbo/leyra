package model

type User struct {
	Name string
	Age  int
}

var Users []User

func NewUser() User {
	return User{}
}

func (u User) Save() bool {
	Users = append(Users, u)
	return true
}
