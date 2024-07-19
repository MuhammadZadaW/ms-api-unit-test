package test

import (
	"ms-api-unit-test/controller"
	"ms-api-unit-test/model"
	"ms-api-unit-test/usecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFindAll(t *testing.T) {
	r := gin.Default()
	api := r.Group("/api")

	userRepoMock := CreateUserRepoMock(mock.Mock{})
	userUsecase := usecase.CreateUserUsecase(userRepoMock)
	controller.CreateUserController(api, userUsecase)

	mockUsers := []model.User{
		{ID: 1, FirstName: "John", LastName: "Doe", Email: "5L7wK@example.com"},
		{ID: 2, FirstName: "Jane", LastName: "Doe", Email: "5L7wK@example.com"},
	}
	
	temp, _ := userRepoMock.(*UserRepoStructMock)
	temp.Mock.On("FindAll").Return(mockUsers)

	req, _ := http.NewRequest("GET", "/api/users", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, ObjectToString(mockUsers), w.Body.String())
	temp.Mock.AssertExpectations(t)
}

func TestFindAllEmpty(t *testing.T) {
	r := gin.Default()
	api := r.Group("/api")

	userRepoMock := CreateUserRepoMock(mock.Mock{})
	userUsecase := usecase.CreateUserUsecase(userRepoMock)
	controller.CreateUserController(api, userUsecase)

	temp, _ := userRepoMock.(*UserRepoStructMock)
	temp.Mock.On("FindAll").Return([]model.User{})

	req, _ := http.NewRequest("GET", "/api/users", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, ObjectToString([]model.User{}), w.Body.String())
	temp.Mock.AssertExpectations(t)
}

func TestGetById(t *testing.T) {
	r := gin.Default()
	api := r.Group("/api")

	userRepoMock := CreateUserRepoMock(mock.Mock{})
	userUsecase := usecase.CreateUserUsecase(userRepoMock)
	controller.CreateUserController(api, userUsecase)

	mockUser := model.User{ID: 1, FirstName: "John", LastName: "Doe", Email: "5L7wK@example.com"}

	temp, _ := userRepoMock.(*UserRepoStructMock)
	temp.Mock.On("GetById", 1).Return(mockUser)

	req, _ := http.NewRequest("GET", "/api/user/1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, ObjectToString(mockUser), w.Body.String())
	temp.Mock.AssertExpectations(t)
}

func TestGetByIdNotFound(t *testing.T) {
	r := gin.Default()
	api := r.Group("/api")

	userRepoMock := CreateUserRepoMock(mock.Mock{})
	userUsecase := usecase.CreateUserUsecase(userRepoMock)
	controller.CreateUserController(api, userUsecase)

	temp, _ := userRepoMock.(*UserRepoStructMock)
	temp.Mock.On("GetById", 1).Return(model.User{})

	req, _ := http.NewRequest("GET", "/api/user/1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, ObjectToString(struct{
		Message string `json:"message"`
	}{
		Message: "user not found",
	}), w.Body.String())
	temp.Mock.AssertExpectations(t)
}


// Benchmark
func BenchmarkFindAll(b *testing.B) {
	r := gin.Default()
	api := r.Group("/api")

	userRepoMock := CreateUserRepoMock(mock.Mock{})
	userUsecase := usecase.CreateUserUsecase(userRepoMock)
	controller.CreateUserController(api, userUsecase)

	temp, _ := userRepoMock.(*UserRepoStructMock)
	temp.Mock.On("FindAll").Return([]model.User{
		{ID: 1, FirstName: "John", LastName: "Doe", Email: "5L7wK@example.com"},
		{ID: 2, FirstName: "Jane", LastName: "Doe", Email: "5L7wK@example.com"},
	})

	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("GET", "/api/users", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			b.Fail()
		}
	}
}

func BenchmarkGetById(b *testing.B) {
	r := gin.Default()
	api := r.Group("/api")

	userRepoMock := CreateUserRepoMock(mock.Mock{})
	userUsecase := usecase.CreateUserUsecase(userRepoMock)
	controller.CreateUserController(api, userUsecase)

	temp, _ := userRepoMock.(*UserRepoStructMock)
	temp.Mock.On("GetById", 1).Return(model.User{
		ID: 1, FirstName: "John", LastName: "Doe", Email: "5L7wK@example.com",
	})

	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("GET", "/api/user/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			b.Fail()
		}
	}
}
