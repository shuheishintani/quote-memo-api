package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shuheishintani/quote-memo-api/src/models"
)

func (ctl *Controller) GetPublicQuotes(c *gin.Context) {
	strTags := c.Query("tags")
	var tagNames []string
	if strTags != "" {
		tagNames = strings.Split(strTags, ",")
	}

	page := c.Query("page")
	i, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	quotes, err := ctl.service.GetPublicQuotes(tagNames, 10*(i-1), 10)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quotes)
}

func (ctl *Controller) GetPrivateQuotes(c *gin.Context) {
	uid := c.GetString("uid")
	strTags := c.Query("tags")
	var tagNames []string
	if strTags != "" {
		tagNames = strings.Split(strTags, ",")
	}

	page := c.Query("page")
	i, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	quotes, err := ctl.service.GetPrivateQuotes(tagNames, uid, 10*(i-1), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quotes)
}

func (ctl *Controller) GetFavoriteQuotes(c *gin.Context) {
	uid := c.GetString("uid")

	quotes, err := ctl.service.GetFavoriteQuotes(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quotes)
}

func (ctl *Controller) GetPrivateQuotesForExport(c *gin.Context) {
	uid := c.GetString("uid")

	quotes, err := ctl.service.GetPrivateQuotesForExport(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quotes)
}

func (ctl *Controller) PostQuote(c *gin.Context) {
	uid := c.GetString("uid")

	postQuoteInput := models.Quote{}
	if err := c.BindJSON(&postQuoteInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctl.validator.Struct(postQuoteInput); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	quote, err := ctl.service.PostQuote(postQuoteInput, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, quote)
}

func (ctl *Controller) UpdateQuote(c *gin.Context) {
	uid := c.GetString("uid")
	id := c.Param("id")

	quote, err := ctl.service.GetQuoteById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if quote.UserID != uid {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden action"})
		return
	}

	updateQuoteInput := models.Quote{}
	if err := c.BindJSON(&updateQuoteInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := ctl.service.UpdateQuote(updateQuoteInput, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (ctl *Controller) DeleteQuote(c *gin.Context) {
	uid := c.GetString("uid")
	id := c.Param("id")

	quote, err := ctl.service.GetQuoteById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if quote.UserID != uid {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden action"})
		return
	}

	result, err := ctl.service.DeleteQuote(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (ctl *Controller) AddFavoriteQuote(c *gin.Context) {
	uid := c.GetString("uid")
	id := c.Param("id")

	user, err := ctl.service.AddFavoriteQuote(uid, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (ctl *Controller) RemoveFavoriteQuote(c *gin.Context) {
	uid := c.GetString("uid")
	id := c.Param("id")

	user, err := ctl.service.RemoveFavoriteQuote(uid, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
