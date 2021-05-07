package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) GetTags(c *gin.Context) {
	tags, err := ctl.service.GetTags()
	if err != nil {
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}
	c.JSON(http.StatusOK, tags)
}
