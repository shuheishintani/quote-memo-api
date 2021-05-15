package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shuheishintani/quote-manager-api/src/dto"
)

func (ctl *Controller) CreateOrUpdateUser(c *gin.Context) {
	uid := c.GetString("uid")
	fmt.Println(uid)
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
