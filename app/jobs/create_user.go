package job

import (
	"leyra/app/models"
)

type CreateUser struct {
	name string
	age  int
}

func NewCreateUser(name string, age int) CreateUser {
	return CreateUser{
		name: name,
		age:  age,
	}
}

func (cu CreateUser) Handle() bool {
	m := model.NewUser()
	m.Name = cu.name
	m.Age = cu.age

	return m.Save()
}
