package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shuheishintani/quote-manager-api/src/dto"
)

func (ctl *Controller) GetPrivateQuotes(c *gin.Context) {
	uid := c.GetString("uid")

	strTags := c.Query("tags")
	var tagNames []string
	if strTags != "" {
		tagNames = strings.Split(strTags, ",")
	}

	quotes, err := ctl.service.GetPrivateQuotes(tagNames, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quotes)
}

func (ctl *Controller) PostQuote(c *gin.Context) {
	uid := c.GetString("uid")

	postQuoteInput := dto.PostQuoteInput{}
	if err := c.BindJSON(&postQuoteInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	quote, err := ctl.service.PostQuote(postQuoteInput, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quote)
}
