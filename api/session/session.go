package session

import (
	"net/http"

	"github.com/da4nik/swanager/api/common"
	"github.com/da4nik/swanager/core/auth"
	"github.com/gin-gonic/gin"
)

type loginMessage struct {
	Email      string
	Password   string
	RememberMe bool `json:"remember_me,omitempty"`
}

// GetRoutesForRouter adds resource routes to api router
func GetRoutesForRouter(router *gin.RouterGroup) *gin.RouterGroup {

	auth := router.Group("/session")
	{
		auth.POST("", common.Auth(false), login)
		auth.DELETE("", common.Auth(true), logout)
	}

	return auth
}

func login(c *gin.Context) {
	var json loginMessage
	if err := c.BindJSON(&json); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := auth.WithEmailAndPassword(json.Email, json.Password)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func logout(c *gin.Context) {
	user := common.MustGetCurrentUser(c)
	if err := auth.Deauthorize(user); err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}
	c.AbortWithStatus(http.StatusOK)
}
