package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) GetTags(c *gin.Context) {
	tags, err := ctl.service.GetTags()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tags)
}
