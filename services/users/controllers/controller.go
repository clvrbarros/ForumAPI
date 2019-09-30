package controllers

import (
	"github.com/clvrbarros/ForumAPI/services/users/models"
	"gopkg.in/go-playground/validator.v9"
)

// Setup for controllers package
type Setup struct {
	model *models.Conn
	validator *validator.Validate
}

// Init the setup
func NewController(model *models.Conn, validator *validator.Validate) *Setup {
	return &Setup{model, validator}
}