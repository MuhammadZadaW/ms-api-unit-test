package main

import (
	"ms-api-unit-test/controller"
	"ms-api-unit-test/usecase"
	"ms-api-unit-test/repo"
	"github.com/gin-gonic/gin"
)

func main() {
	
	r := gin.Default()
	api := r.Group("/api")

	repoUser := repo.CreateUserRepo()
	usecaseUser := usecase.CreateUserUsecase(repoUser)
	controller.CreateUserController(api, usecaseUser)
	
	r.Run()
}