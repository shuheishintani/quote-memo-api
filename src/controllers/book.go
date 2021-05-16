package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shuheishintani/quote-manager-api/src/services"
)

func (ctl *Controller) GetBooks(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")
	page := c.Query("page")
	getBooksInput := services.GetBooksQuery{
		Title:  title,
		Author: author,
		Page:   page,
	}

	books, err := ctl.service.GetBooks(getBooksInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}
