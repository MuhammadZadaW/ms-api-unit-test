package repo

import (
	"ms-api-unit-test/model"
)

var DataList = []model.User{
	{ID: 1, FirstName: "John", LastName: "Doe", Email: "5L7wK@example.com"},
	{ID: 2, FirstName: "Jane", LastName: "Doe", Email: "5L7wK@example.com"},
}

type UserRepoStruct struct{}

type UserRepoIface interface {
	FindAll() []model.User
	GetById(id int) model.User
}

func CreateUserRepo() UserRepoIface {
	return &UserRepoStruct{}
}

func (u *UserRepoStruct) FindAll() []model.User {
	return DataList
}

func (u *UserRepoStruct) GetById(id int) model.User {

	for _, data := range DataList {
		if data.ID == id {
			return data
		}
	}
	
	return model.User{}
}