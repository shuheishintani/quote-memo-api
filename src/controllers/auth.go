package controllers

import (
	"context"
	"net/http"
	"time"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

func (ctl *Controller) Login(c *gin.Context) {
	auth := c.MustGet("auth").(*auth.Client)
	type RequestParams struct {
		IDToken string `json:"id_token"`
	}
	params := RequestParams{}

	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expiresIn := time.Hour * 24 * 5

	cookie, err := auth.SessionCookie(context.Background(), params.IDToken, expiresIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("session", cookie, int(expiresIn.Seconds()), "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (ctl *Controller) Logout(c *gin.Context) {
	c.SetCookie("session", "", 0, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
