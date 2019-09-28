package controllers
import (
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
}