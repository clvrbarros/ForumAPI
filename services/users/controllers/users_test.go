package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	//body := gin.H{
	//	"email": "email@gmail.com",
	//	"firstname": "name",
	//	"password": "123456",
	//}
	r := gin.New()
	req, _  := http.NewRequest("POST","/register", nil)

	w := httptest.NewRecorder()
	r.Handle("POST", "/register", CreateUser)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}