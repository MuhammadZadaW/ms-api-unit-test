package usecase

import (
	"ms-api-unit-test/model"
	"ms-api-unit-test/repo"
)

type UserUsecaseStruct struct {
	userRepo repo.UserRepoIface
}

type UserUsecaseIface interface {
	FindAll() []model.User
	GetById(id int) model.User
}

func CreateUserUsecase(userRepo repo.UserRepoIface) UserUsecaseIface {
	return &UserUsecaseStruct{userRepo}
}

func (u *UserUsecaseStruct) FindAll() []model.User {
	return u.userRepo.FindAll()
}

func (u *UserUsecaseStruct) GetById(id int) model.User {
	return u.userRepo.GetById(id)
}