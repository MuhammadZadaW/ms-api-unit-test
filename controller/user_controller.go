package controller

import (
	"strconv"
	"ms-api-unit-test/model"
	"ms-api-unit-test/usecase"
	"github.com/gin-gonic/gin"
)

type UserControllerStruct struct {
	userUsecase usecase.UserUsecaseIface
}

func CreateUserController(r *gin.RouterGroup, userUsecase usecase.UserUsecaseIface) {
	controller := &UserControllerStruct{userUsecase}

	r.GET("/users", controller.FindAll)
	r.GET("/user/:id", controller.GetById)
}

func (u *UserControllerStruct) FindAll(c *gin.Context) {
	users := u.userUsecase.FindAll()
	c.JSON(200, users)
}

func (u *UserControllerStruct) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := u.userUsecase.GetById(id)

	if user == (model.User{}) {
		c.JSON(404, gin.H{"message": "user not found"})
		return
	}

	c.JSON(200, user)
}