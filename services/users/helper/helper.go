package helper

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strings"
	"golang.org/x/crypto/bcrypt"
)

// APIMessage type is a struct for generic JSON response
type APIMessage struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
}

// APIValidator type is a struct for multiple error messages
type APIValidator struct {
	Errors []APIMessage `json:"errors"`
}

// HashPassword encrypts a given password using bcrypt algorithm
func HashPassword(password string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// CheckPasswordHash checks if the given password matches
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ValidatorMessage validate error messages
func ValidatorMessage(c *gin.Context, err error) {
	apiValidator := &APIValidator{}

	for _, err := range err.(validator.ValidationErrors) {
		var customError string

		switch err.Tag() {
		case "required":
			customError = "is required"
		case "email":
			customError = "is not valid"
		case "min":
			customError = "minimum length is " + err.Param()
		}
		message := APIMessage{
			Message: "The " + strings.ToLower(err.Field()) + " field " + customError,
		}
		apiValidator.Errors = append(apiValidator.Errors, message)
	}
	c.JSON(http.StatusBadRequest, apiValidator)
}