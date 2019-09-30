package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/clvrbarros/ForumAPI/services/users/helper"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestCreateUser(t *testing.T) {
	const (
		method = "POST"
		path = "/register"
	)
	router.Handle(method, path, h.CreateUser)

	t.Run("Should return 201 when successful", func(t *testing.T) {
		var newUser = []byte(`{
				"email": "teste@gmail.com",
				"firstname": "Test",
				"password": "123456"
				}`)

		req, rr := request(method, path, bytes.NewBuffer(newUser), t)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
	})

	t.Run("Should return 400 and error message", func(t *testing.T) {
		t.Run("when body is empty", func(t *testing.T) {
			req, rr := request(method, path, nil, t)
			router.ServeHTTP(rr, req)

			assert.Equal(t, http.StatusBadRequest, rr.Code)
		})
		t.Run("when email is not valid", func(t *testing.T) {
			var newUser = []byte(`{
				"email": "testegmail.com",
				"firstname": "Test",
				"password": "1236"
				}`)
			var response helper.APIValidator

			req, rr := request(method, path, bytes.NewBuffer(newUser), t)
			router.ServeHTTP(rr, req)

			readBuf, _ := ioutil.ReadAll(rr.Body)
			err := json.Unmarshal(readBuf, &response)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t,"The email field is not valid", response.Errors[0].Message)
			assert.Equal(t, http.StatusBadRequest, rr.Code)
		})
		t.Run("when firstname is empty", func(t *testing.T) {
			var newUser = []byte(`{
				"email": "teste@gmail.com",
				"firstname": "",
				"password": "123645"
				}`)
			var response helper.APIValidator

			req, rr := request(method, path, bytes.NewBuffer(newUser), t)
			router.ServeHTTP(rr, req)

			readBuf, _ := ioutil.ReadAll(rr.Body)
			err := json.Unmarshal(readBuf, &response)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t,"The firstname field is required", response.Errors[0].Message)
			assert.Equal(t, http.StatusBadRequest, rr.Code)
		})
	})
}