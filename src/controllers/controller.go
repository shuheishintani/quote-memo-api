package controllers

import (
	"github.com/go-playground/validator"
	"github.com/shuheishintani/quote-memo-api/src/services"
)

type Controller struct {
	service   *services.Service
	validator *validator.Validate
}

func NewController(service *services.Service) *Controller {
	return &Controller{
		service:   service,
		validator: validator.New(),
	}
}
