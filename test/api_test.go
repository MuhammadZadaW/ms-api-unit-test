package test

import (
	"ms-api-unit-test/controller"
	"ms-api-unit-test/usecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"ms-api-unit-test/repo"
	"io/ioutil"
	"encoding/json"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	gin.DefaultWriter = ioutil.Discard
	gin.DefaultErrorWriter = ioutil.Discard
	m.Run()
}

func TestApiFindAll(t *testing.T) {
	r := gin.Default()
	api := r.Group("/api")

	userRepo := repo.CreateUserRepo()
	userUsecase := usecase.CreateUserUsecase(userRepo)
	controller.CreateUserController(api, userUsecase)

	req, _ := http.NewRequest("GET", "/api/users", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, ObjectToString(repo.DataList), w.Body.String())
}

func TestApiGetById(t *testing.T) {
	r := gin.Default()
	api := r.Group("/api")

	userRepo := repo.CreateUserRepo()
	userUsecase := usecase.CreateUserUsecase(userRepo)
	controller.CreateUserController(api, userUsecase)

	req, _ := http.NewRequest("GET", "/api/user/1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestApiGetByIdNotFound(t *testing.T) {
	r := gin.Default()
	api := r.Group("/api")

	userRepo := repo.CreateUserRepo()
	userUsecase := usecase.CreateUserUsecase(userRepo)
	controller.CreateUserController(api, userUsecase)

	req, _ := http.NewRequest("GET", "/api/user/0", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, ObjectToString(struct{
		Message string `json:"message"`
	}{
		Message: "user not found",
	}), w.Body.String())
}

func ObjectToString(o interface{}) string {
	out, _ := json.Marshal(o)
	return string(out)
}

