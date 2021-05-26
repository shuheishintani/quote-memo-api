package controllers

import (
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/shuheishintani/quote-memo-api/src/models"
)

func (ctl *Controller) CreateOrUpdateUser(c *gin.Context) {
	userInput := models.User{}
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

func (ctl *Controller) GetUsers(c *gin.Context) {
	users, err := ctl.service.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (ctl *Controller) DeleteUser(c *gin.Context) {
	auth := c.MustGet("auth").(*auth.Client)
	uid := c.GetString("uid")
	_, err := ctl.service.GetUserById(uid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	result, err := ctl.service.DeleteUser(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := auth.DeleteUser(c, uid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, result)
}
