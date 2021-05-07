package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shuheishintani/quote-manager-api/src/dto"
	"github.com/shuheishintani/quote-manager-api/src/services"
)

type Controller struct {
	service *services.Service
}

func NewController(service *services.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctl *Controller) GetQuotes(c *gin.Context) {
	// uid := c.MustGet("uid").(string)
	// println(uid)
	strTags := c.Query("tags")
	var tagNames []string
	if strTags != "" {
		tagNames = strings.Split(strTags, ",")
	}
	quotes, err := ctl.service.GetQuotes(tagNames)
	if err != nil {
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}
	c.JSON(http.StatusOK, quotes)
}

func (ctl *Controller) PostQuote(c *gin.Context) {
	postQuoteInput := dto.PostQuoteInput{}
	if err := c.BindJSON(&postQuoteInput); err != nil {
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}

	fmt.Println(postQuoteInput)

	quote, err := ctl.service.PostQuote(postQuoteInput)
	if err != nil {
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}
	c.JSON(http.StatusOK, quote)
}
