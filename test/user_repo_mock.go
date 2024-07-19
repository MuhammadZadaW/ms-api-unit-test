package test

import (
	"ms-api-unit-test/model"
	"ms-api-unit-test/repo"
	"github.com/stretchr/testify/mock"
)

type UserRepoStructMock struct {
	Mock mock.Mock
}

func CreateUserRepoMock(mock mock.Mock) repo.UserRepoIface {
	return &UserRepoStructMock{
		Mock: mock,
	}
}

func (u *UserRepoStructMock) FindAll() []model.User {
	
	args := u.Mock.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]model.User)
	}

	return []model.User{}
}

func (u *UserRepoStructMock) GetById(id int) model.User {

	args := u.Mock.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(model.User)
	} 

	return model.User{}
}