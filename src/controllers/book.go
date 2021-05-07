package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shuheishintani/quote-manager-api/src/dto"
)

func (ctl *Controller) GetBooks(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")
	page := c.Query("page")

	getBooksInput := dto.GetBooksInput{
		Title:  title,
		Author: author,
		Page:   page,
	}
	books, err := ctl.service.GetBooks(getBooksInput)
	if err != nil {
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}
	c.JSON(http.StatusOK, books)
}
