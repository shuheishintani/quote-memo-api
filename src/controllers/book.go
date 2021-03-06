package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shuheishintani/quote-memo-api/src/services"
)

func (ctl *Controller) GetExternalBooks(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")
	page := c.Query("page")
	getBooksInput := services.GetBooksQuery{
		Title:  title,
		Author: author,
		Page:   page,
	}

	books, err := ctl.service.GetExternalBooks(getBooksInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (ctl *Controller) GetBooks(c *gin.Context) {
	keyword := c.Query("keyword")
	books, err := ctl.service.GetBooks(keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (ctl *Controller) GetBookById(c *gin.Context) {
	id := c.Param("id")
	book, err := ctl.service.GetBookById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}
