package controllers

import "github.com/shuheishintani/quote-manager-api/src/services"

type Controller struct {
	service *services.Service
}

func NewController(service *services.Service) *Controller {
	return &Controller{
		service: service,
	}
}
