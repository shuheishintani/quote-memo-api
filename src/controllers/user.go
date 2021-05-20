package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shuheishintani/quote-memo-api/src/dto"
)

func (ctl *Controller) CreateOrUpdateUser(c *gin.Context) {
	userInput := dto.UserInput{}
	if err := c.BindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tags, err := ctl.service.CreateOrUpdateUser(userInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tags)
}

func (ctl *Controller) GetMe(c *gin.Context) {
	uid := c.GetString("uid")
	user, err := ctl.service.GetUserById(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (ctl *Controller) GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := ctl.service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
